package main

import "fmt"

func main() {
	s := []int{1, 2, 3}

	sum := 0
	for _, v := range s {
		sum += v
	}
	fmt.Println(sum)
}
