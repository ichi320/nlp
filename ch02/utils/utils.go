package utils

import (
	"fmt"
	"log"
	"math"
	"strings"

	"gonum.org/v1/gonum/mat"
)

// Preprocess 単語を数値にしたコーパスとその変換辞書を作成する
func Preprocess(text string) ([]int, map[string]int, map[int]string) {
	text = strings.ToLower(text)
	text = strings.Replace(text, ".", " .", -1)
	var words []string
	words = strings.Split(text, " ")

	wordToID := map[string]int{}
	idToWord := map[int]string{}
	for _, word := range words {
		_, ok := wordToID[word]
		if !ok {
			newID := len(wordToID)
			wordToID[word] = newID
			idToWord[newID] = word
		}
	}

	var corpus []int
	for _, v := range words {
		corpus = append(corpus, wordToID[v])
	}
	return corpus, wordToID, idToWord
}

// CreateCoMatrix 共起行列を作る
func CreateCoMatrix(corpus []int, vocabSize int, windowSize int) *mat.Dense {
	corpusSize := len(corpus)
	var coMatrix *mat.Dense
	coMatrix = mat.NewDense(vocabSize, vocabSize, nil)
	for idx, wordID := range corpus {
		for i := 1; i <= windowSize; i++ {
			leftIdx := idx - i
			rightIdx := idx + i
			if leftIdx >= 0 {
				leftWordID := corpus[leftIdx]
				tmpValue := coMatrix.At(wordID, leftWordID)
				tmpValue++
				coMatrix.Set(wordID, leftWordID, tmpValue)
			}
			if rightIdx < corpusSize {
				rightWordID := corpus[rightIdx]
				tmpValue := coMatrix.At(wordID, rightWordID)
				tmpValue++
				coMatrix.Set(wordID, rightWordID, tmpValue)
			}
		}

	}

	return coMatrix
}

// CosSimilarity コサイン類似度を計算する
func CosSimilarity(a, b mat.Vector) float64 {
	eps := 1e-8
	var norma, normb float64
	norma = mat.Dot(a, a)
	norma = math.Sqrt(norma) + eps
	normb = mat.Dot(b, b)
	normb = math.Sqrt(normb) + eps

	sim := mat.Dot(a, b) / (norma * normb)
	return math.Round(sim*1000) / 1000
}

// MostSimilar (query string, wordToID []int, idToWord []int, wordMatrix mat.Matrix) {
func MostSimilar(query string, wordToID map[string]int, idToWord map[int]string, wordMatrix *mat.Dense, top int) {
	queryID, ok := wordToID[query]
	if !ok {
		fmt.Printf("%v is not found.", query)
		return
	}

	queryVec := wordMatrix.RowView(queryID)
	vocabSize := len(wordToID)
	similarity := make([]float64, vocabSize)
	for i := 0; i < vocabSize; i++ {
		similarity[i] = CosSimilarity(wordMatrix.RowView(i), queryVec)
	}

	count := 0
	for i := 0; i < len(similarity); i++ {
		similarity[i] *= -1
	}
	rank := ArgsortNew(similarity)
	for i := 0; i < len(similarity); i++ {
		similarity[i] *= -1
	}
	for _, i := range rank {
		if idToWord[i] == query {
			continue
		}
		fmt.Printf("%v: %v\n", idToWord[i], similarity[i])
		count++
		if count >= top {
			return
		}
	}
}

// Ppmi calc pmi
func Ppmi(C *mat.Dense) *mat.Dense {
	verbose := true
	eps := 1e-8
	r, c := C.Dims()
	M := mat.NewDense(r, c, nil)
	var sum float64
	sum = 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			sum += C.At(i, j)
		}
	}
	N := sum
	var rawS []float64
	for i := 0; i < r; i++ {
		sum = 0
		for j := 0; j < c; j++ {
			sum += C.At(i, j)
		}
		rawS = append(rawS, sum)
	}
	var S *mat.VecDense
	S = mat.NewVecDense(len(rawS), rawS)
	total := r * c
	cnt := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			pmi := math.Max(0, math.Log2(C.At(i, j)*N/(S.At(j, 0)*S.At(i, 0))+eps))
			pmi = math.Round(pmi*1000) / 1000
			M.Set(i, j, pmi)
			if verbose {
				cnt++
				if math.Mod(float64(cnt), float64(math.Floor(float64(total)/10))) == 0 {
					fmt.Printf("%.2f done.\n", float64(cnt)/float64(total))
				}
			}
		}
	}
	return M
}

// CalcSVDU calc U from matrix
func CalcSVDU(a *mat.Dense) *mat.Dense {

	var svd mat.SVD
	ok := svd.Factorize(a, mat.SVDFull)
	if !ok {
		log.Fatal("failed to factorize A")
	}

	var c *mat.Dense
	c = &mat.Dense{}
	svd.UTo(c)

	return c

}
