package main

import (
	"fmt"
	"math/big"
)

func main() {
	// 大整数
	fmt.Printf("big.NewInt(123): %v\n", big.NewInt(123))
	// 大有理数
	b := big.NewRat(1,3)
	b.Add(big.NewRat(1,3),big.NewRat(1,3))
	fmt.Printf("b: %v\n", b)
	b.Add(big.NewRat(1,3),big.NewRat(1,3))
	fmt.Printf("b: %v\n", b)

	i := big.NewInt(123)
	i.Add(big.NewInt(14),big.NewInt(15))
	fmt.Printf("i: %v\n", i)
}
