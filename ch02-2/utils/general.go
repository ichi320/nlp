package utils

import (
	"fmt"
	"os"

	"gonum.org/v1/gonum/mat"
)

// MatPrint print out Matrix
func MatPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
	fmt.Println("")
}

// MatPrintRound print out Matrix rounded
func MatPrintRound(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%.3f\n", fa)
	fmt.Println("")
}

// NpSum calc sum of vector
func NpSum(x []float64) float64 {
	var sum float64
	for _, v := range x {
		sum += v
	}
	return sum
}

// NpSquareSum calc sum of vector(x**2)
func NpSquareSum(x []float64) float64 {
	var x2 []float64
	for _, v := range x {
		x2 = append(x2, v*v)
	}
	return NpSum(x2)
}

// CreateDataDir send data directory
// in: none, out: string, error
func CreateDataDir() (string, error) {
	dataDir, _ := os.Getwd() // get current directory
	if d, err := os.Stat(dataDir + "/source"); os.IsNotExist(err) || !d.IsDir() {
		if err := os.Mkdir(dataDir+"/source", 0777); err != nil {
			return "", err
		}
	}
	return dataDir + "/source", nil

}
