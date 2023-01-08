package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

func main() {
	var u U
	file, _ := os.Open("gob.txt")

	decode := gob.NewDecoder(file)
	decode.Decode(&u)
	fmt.Printf("u: %v\n", u) // 对面Y是零值没有传递
}

type U struct {
	X, Y *int8
}