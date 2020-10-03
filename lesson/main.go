package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func matPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}

func main() {
	fmt.Println("Hello, world!")
	A := mat.NewDense(3, 4, nil)
	fmt.Println(A)
	matPrint(A)

	x := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	B := mat.NewDense(3, 4, x)
	matPrint(B)

	b := B.At(0, 2)
	fmt.Println(b)

	B.Set(0, 2, -1.5)
	matPrint(B)

	matPrint(B.RowView(1))

	col := []float64{3, 2, 1}
	B.SetCol(2, col)
	matPrint(B)
}
