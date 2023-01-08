package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	sum := sum(arr[:])
	fmt.Printf("sum: %v\n", sum)
}

func sum(s []int) (sum int) {
	for _, v := range s {
		sum += v
	}
	return
}
