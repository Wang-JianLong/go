package main

import "fmt"

func main() {
	// 增加切片长度
	j := 5
	s := []int{1,2,3}
	s = append(s, make([]int, j)...)
	fmt.Printf("len(s): %v\n", len(s))
	// 在索引 i 1的位置插入元素 x(9)
	a := []int{1,2,3}
	i := 1
	a = append(a[:i],append([]int{9},a[i:]...)...) // a: [1 9 2 3]
	fmt.Printf("a: %v\n", a)
	// 在索引 i 的位置插入长度为 j 的新切片
	i =  1
	s1 := []int{1,2,3}

	s1 = append(s1[:i],append(make([]int,5),s1[i:]...)...) // s1: [1 0 0 0 0 0 2 3]
	fmt.Printf("s1: %v\n", s1)
	// 在索引 i 的位置插入切片 b 的所有元素
	s2 := []int{1,2,3,4}
	i = 2
	newS := []int{99,100}

	s2 = append(s2[:i],append(newS,s[i:]...)...)
	fmt.Printf("s2: %v\n", s2)

	
}