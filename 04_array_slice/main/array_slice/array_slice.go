package main

import "fmt"

func main() {
	sl := [7]int{1, 2, 3, 4, 5, 6, 7, 8}

	s1 := sl[:]

	s1[7] = 100 // error 数组越界异常

	fmt.Printf("s1: %v\n", s1)
}