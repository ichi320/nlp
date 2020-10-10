package main

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/mat"
)

func matPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
	fmt.Println("")
}

// Add 感覚に合わないので、関数を作る
func Add(a mat.Matrix, b mat.Matrix) mat.Matrix {
	// var B mat.Dense
	// fmt.Printf("B type : %T\n%v\n", B, B)
	// B.Add(a, b)
	// return &B
	r, c := a.Dims()
	B := mat.NewDense(r, c, nil)
	fmt.Printf("B type : %T\n%v\n", B, B)
	B.Add(a, b)
	return B
}

// Sigmoid is in: Matrix out:Matrix
func Sigmoid(x mat.Matrix) mat.Matrix {
	sigmoid := func(i, j int, v float64) float64 {
		return 1 / (1 + math.Exp(-v))
	}
	var result *mat.Dense
	r, c := x.Dims()
	result = mat.NewDense(r, c, nil) // この空き箱を作らないとエラーになる
	result.Apply(sigmoid, x)
	return result
}

func main() {
	fmt.Println("Hello, world!")
	A := mat.NewDense(3, 4, nil)
	fmt.Printf("%T\n", A)
	matPrint(A)

	x := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	B := mat.NewDense(3, 4, x)
	matPrint(B)

	// b := B.At(0, 2)
	// fmt.Println(b)

	// B.Set(0, 2, -1.5)
	// matPrint(B)

	// fmt.Printf("%T\n", B.RowView(1)) // *mat.VecDense
	// matPrint(B.RowView(1))

	// col := []float64{3, 2, 1}
	// B.SetCol(2, col)
	// matPrint(B)

	// // add
	// C := mat.NewDense(3, 4, nil)
	// C.Add(B, B)
	// matPrint(C)

	// // sub
	// C.Sub(B, B)
	// matPrint(C)

	// 作った関数にて
	var C2 mat.Matrix
	C2 = Add(B, B)
	matPrint(C2)

	// // scale 定数倍
	// C = mat.NewDense(3, 4, nil)
	// C.Scale(5, B)
	// matPrint(C)

	// // 転置
	// D := B.T()
	// matPrint(D)

	// // 内積
	// matPrint(B)
	// E := mat.NewDense(3, 3, nil)
	// E.Product(B, B.T())
	// matPrint(E)

	// // 要素積
	// matPrint(B)
	// B.MulElem(B, B)
	// matPrint(B)

	// // スライス
	// F := B.Slice(0, 2, 0, 2)
	// matPrint(F)

	// // 行列の縦の結合はStack、横の結合はAugment
	// G := mat.NewDense(3, 2, []float64{
	// 	1, 2,
	// 	3, 4,
	// 	5, 6,
	// })
	// H := mat.NewDense(3, 2, []float64{
	// 	7, 8,
	// 	9, 10,
	// 	11, 12,
	// })

	// I := mat.NewDense(6, 2, nil)
	// I.Stack(G, H)
	// matPrint(I)

	// J := mat.NewDense(3, 4, nil)
	// J.Augment(G, H)
	// matPrint(J)

	// Apply
	sumOfIndices := func(i, j int, v float64) float64 {
		return float64(i+j) + v
	}

	matPrint(B)
	var K *mat.Dense
	r, c := B.Dims()
	K = mat.NewDense(r, c, nil)
	K.Apply(sumOfIndices, B)
	matPrint(K)

	L := Sigmoid(B)
	matPrint(L)

}
