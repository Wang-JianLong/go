package main

import "fmt"

func main() {
	var a1 [3][5]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			a1[i][j] = i + j
		}
	}

	fmt.Printf("a1: %v\n", a1) // a1: [[0 1 2 3 4] [1 2 3 4 5] [2 3 4 5 6]]

}