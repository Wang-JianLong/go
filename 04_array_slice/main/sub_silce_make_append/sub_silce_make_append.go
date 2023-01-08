package main

import "fmt"

func main() {
	s := make([]int, 3, 5) // len3 cap5

	sub := s[2:] // len1 cap3

	sub = append(sub, 1)
	sub[0] = 1
	fmt.Printf("sub: %v\n", sub) // sub: [1 1]
	fmt.Printf("s: %v\n", s)     // s: [0 0 1]
	// 如果 此时sub的数未超过s的cap那么还是同一块内存
	sub = append(sub, 1)
	sub[0] = 100
	fmt.Printf("sub: %v\n", sub) // sub: [100 1 1]
	fmt.Printf("s: %v\n", s)     // s: [0 0 100]

	// 此时再append sub一个元素 那么 从sub起始的位置开始 2 + 2 = 4  下标为4 实际5个元素  append再来一个就是6个 那么此时会怎么样？

	sub = append(sub, 1)
	sub[0] = 1000
	fmt.Printf("sub: %v\n", sub) // sub: [1000 1 1 1]
	fmt.Printf("s: %v\n", s)     // s: [0 0 100]
	// 此时 由于sub超过了母切片的cap那么 sub就会重新分配一份内存空间使用

}
