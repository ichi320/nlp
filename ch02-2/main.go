package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"

	"github.com/ichi320/nlp/ch02-2/utils"
)

func main() {
	dataDir, _ := utils.CreateDataDir()
	corpus, wordToID, idToWord := utils.LoadData(dataDir, "train")
	fmt.Println("counting co-occurrence...")
	C := utils.CreateCoMatrix(corpus, len(wordToID), 1)
	fmt.Println("calculating PPMI...")
	W := utils.Ppmi(C)
	fmt.Println("calculating SVD...")
	U := utils.CalcSVDU(W)

	r, _ := U.Dims()
	wordvecSize := 100
	wordVecs := U.Slice(0, r, 0, wordvecSize).(*mat.Dense)
	querys := []string{"you", "year", "car", "toyota"}
	for _, query := range querys {
		utils.MostSimilar(query, wordToID, idToWord, wordVecs, 5)
	}
	// fmt.Println("corpus size: ", len(corpus))
	// fmt.Println(corpus[:30])
	// fmt.Println(idToWord[0], idToWord[1], idToWord[2])
	// fmt.Println(wordToID["car"], wordToID["happy"], wordToID["lexus"])
}
