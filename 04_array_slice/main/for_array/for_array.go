package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}

	// 1. 根据数组下标遍历
	for i := 0; i < len(array); i++ {
		fmt.Printf("array[i]: %v\n", array[i])
	}
	// 根据 range遍历
	for index,value := range array {
		fmt.Printf("index: %v, value: %v\n", index, value)
	}
}