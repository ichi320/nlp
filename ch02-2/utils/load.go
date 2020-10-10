package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"unsafe"
)

// LoadData load data
func LoadData(dataDir string, dataType string) ([]int, map[string]int, map[int]string) {

	var saveFiles map[string]string = map[string]string{
		"train": "ptb.train.json",
		"test":  "ptb.test.json",
		"valid": "ptb.valid.json",
	}

	var savePath string = dataDir + "/" + saveFiles[dataType]

	wordToID, idToWord := loadVocab(dataDir)
	var corpus []int

	if f, err := os.Stat(savePath); !os.IsNotExist(err) && !f.IsDir() {
		bytes, err := ioutil.ReadFile(savePath)
		if err != nil {
			log.Fatalln(err)
		}

		err = json.Unmarshal(bytes, &corpus)
		if err != nil {
			log.Fatalln(err)
		}
		return corpus, wordToID, idToWord
	}

	filePath, err := downLoadSourceFile(dataDir, dataType)
	if err != nil {
		log.Fatalln(err)
	}
	btext, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalln(err)
	}

	var text string
	text = strings.Replace(*(*string)(unsafe.Pointer(&btext)), "\n", "<eos>", -1)
	text = strings.TrimSpace(text)

	var words []string
	words = strings.Split(text, " ")

	for _, word := range words {
		corpus = append(corpus, wordToID[word])
	}

	bytes, err := json.Marshal(&corpus)
	if err != nil {
		log.Fatalln(err)
	}

	err = ioutil.WriteFile(savePath, bytes, 0644)
	if err != nil {
		log.Fatalln(err)
	}

	return corpus, wordToID, idToWord

}

func loadVocab(dataDir string) (map[string]int, map[int]string) {
	wordToID := map[string]int{}
	idToWord := map[int]string{}

	var vocabFile string = "ptb.wordToID.json"
	var vocabPath string = dataDir + "/" + vocabFile
	if f, err := os.Stat(vocabPath); !os.IsNotExist(err) && !f.IsDir() {
		words, err := ioutil.ReadFile(vocabPath)
		if err != nil {
			log.Fatalln(err)
		}
		err = json.Unmarshal(words, &wordToID)
		if err != nil {
			log.Fatalln(err)
		}

		vocabPath = dataDir + "/ptb.idToWord.json"
		words, err = ioutil.ReadFile(vocabPath)
		if err != nil {
			log.Fatalln(err)
		}
		err = json.Unmarshal(words, &idToWord)
		if err != nil {
			log.Fatalln(err)
		}

		return wordToID, idToWord
	}

	filePath, err := downLoadSourceFile(dataDir, "train")
	if err != nil {
		log.Fatalln(err)
	}

	btext, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalln(err)
	}

	var text string
	text = strings.Replace(*(*string)(unsafe.Pointer(&btext)), "\n", "<eos>", -1)
	text = strings.TrimSpace(text)

	_, wordToID, idToWord = Preprocess(text)
	// TODO
	// wordToID, idToWord を保存する
	bytes, err := json.Marshal(&wordToID)
	if err != nil {
		log.Fatalln(err)
	}
	err = ioutil.WriteFile(vocabPath, bytes, 0644)
	if err != nil {
		log.Fatalln(err)
	}

	vocabFile = "ptb.idToWord.json"
	vocabPath = dataDir + "/" + vocabFile
	bytes, err = json.Marshal(&idToWord)
	if err != nil {
		log.Fatalln(err)
	}
	err = ioutil.WriteFile(vocabPath, bytes, 0644)
	if err != nil {
		log.Fatalln(err)
	}

	return wordToID, idToWord

}

func downLoadSourceFile(dataDir string, dataType string) (string, error) {

	var urlBase string = "https://raw.githubusercontent.com/tomsercu/lstm/master/data"

	var keyFiles map[string]string = map[string]string{
		"train": "ptb.train.txt",
		"test":  "ptb.test.txt",
		"valid": "ptb.valid.txt",
	}

	fileName := keyFiles[dataType]
	fmt.Print(fileName + "ダウンロード中...")

	// TODO for sequence
	filePath := dataDir + "/" + fileName
	if f, err := os.Stat(filePath); !os.IsNotExist(err) && !f.IsDir() {
		fmt.Println("元ファイルはダウンロード済みです")
	}

	urlPath := urlBase + "/" + fileName

	resp, err := http.Get(urlPath)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	out, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}
	fmt.Println("")

	fmt.Println("ダウンロード完了")
	return filePath, nil
}
