package utils

import "sort"

// argsort, like in Numpy, it returns an array of indexes into an array. Note
// that the gonum version of argsort reorders the original array and returns
// indexes to reconstruct the original order.
type argsort struct {
	s    []float64 // Points to orignal array but does NOT alter it.
	inds []int     // Indexes to be returned.
}

func (a argsort) Len() int {
	return len(a.s)
}

func (a argsort) Less(i, j int) bool {
	return a.s[a.inds[i]] < a.s[a.inds[j]]
}

func (a argsort) Swap(i, j int) {
	a.inds[i], a.inds[j] = a.inds[j], a.inds[i]
}

// ArgsortNew allocates and returns an array of indexes into the source float
// array.
func ArgsortNew(src []float64) []int {
	inds := make([]int, len(src))
	for i := range src {
		inds[i] = i
	}
	Argsort(src, inds)
	return inds
}

// Argsort alters a caller-allocated array of indexes into the source float
// array. The indexes must already have values 0...n-1.
func Argsort(src []float64, inds []int) {
	if len(src) != len(inds) {
		panic("floats: length of inds does not match length of slice")
	}
	a := argsort{s: src, inds: inds}
	sort.Sort(a)
}

// https://play.golang.org/p/A3niFJ582wJ
