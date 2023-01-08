package main

import "fmt"

func main() {
	data := []int{1, 2, 3}
	s := make([]int, 3, 6)
	fmt.Printf("s: %v\n", s)
	copy(s, data)
	fmt.Printf("s: %v\n", s)

	data2 := []int{4, 5, 6}

	copy(s[3:6], data2)
	fmt.Printf("s: %v\n", s)
	fmt.Printf("s[3:5]: %v\n", s[3:6])
	/*
	s: [0 0 0]
	s: [1 2 3]
	s: [1 2 3] // 只会打印len的？
	s[3:5]: [4 5 6]
	*/
}
