package utils

import (
	"strings"
)

func Preprocess(text string) ([]int, map[string]int, map[int]string) {
	text = strings.Replace(text, ".", " .", -1)
	var words []string
	words = strings.Split(text, " ")

	wordToID := map[string]int{}
	idToWord := map[int]string{}
	for _, v := range words {
		_, ok := wordToID[v]
		if !ok {
			wordToID[v] = len(wordToID)
			idToWord[len(wordToID)] = v
		}
	}

	var corpus []int
	for _, v := range words {
		corpus = append(corpus, wordToID[v])
	}
	return corpus, wordToID, idToWord
}

func create_co_matrix(corpus []int, vocab_size int, window_size int)
