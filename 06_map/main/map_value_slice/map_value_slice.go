package main

import "fmt"

func main() {
	m1 := make(map[int][]int)

	m1[1] = []int{1, 2, 3}

	m2 := make(map[int]*[]int)

	i := []int{1, 2, 3}
	m2[1] = &i

	fmt.Printf("m1: %v\n", m1)
	fmt.Printf("m2: %v\n", m2)
}