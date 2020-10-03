package main

import (
	"fmt"

	"github.com/ichi320/nlp/ch02/utils"
)

func main() {
	text := "You say goodbye and I say hello."
	corpus, a, b := utils.Preprocess(text)
	fmt.Println(corpus)
	fmt.Println(a)
	fmt.Println(b)
	co_matrix := create_co_matrix(corpus, len(a), 1)
}
