package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	s = getMakeSlice(s, 4)
	fmt.Printf("s: %v\n", s)
	fmt.Printf("len(s): %v\n", len(s))
}

/*
给定 slice s[]int 和一个 int 类型的因子 factor，扩展 s 使其长度为 len(s) * factor。
*/

func getMakeSlice(sl []int, factor int) []int {
	s := make([]int, len(sl)*factor)
	copy(s, sl)
	return s
}
