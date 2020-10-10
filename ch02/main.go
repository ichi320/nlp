package main

import (
	"fmt"

	"github.com/ichi320/nlp/ch02/utils"
)

func main() {
	text := "You say goodbye and I say hello."
	corpus, wordToID, idToWord := utils.Preprocess(text)
	C := utils.CreateCoMatrix(corpus, len(wordToID), 1)
	utils.MatPrint(C)
	c0 := C.RowView(wordToID["you"])
	c1 := C.RowView(wordToID["i"])
	cos := utils.CosSimilarity(c0, c1)
	fmt.Println(cos)
	fmt.Println(corpus)
	fmt.Println(wordToID)
	fmt.Println(idToWord)
	utils.MostSimilar("you", wordToID, idToWord, C, 5)

	W := utils.Ppmi(C)
	utils.MatPrint(W)

	U := utils.CalcSVDU(W)
	utils.MatPrintRound(U)
}
