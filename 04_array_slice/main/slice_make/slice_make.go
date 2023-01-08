package main

import "fmt"

// 当相关数组还没有定义时，我们可以使用 make () 函数来创建一个切片 同时创建好相关数组：var slice1 []type = make([]type, len)。

// 也可以简写为 slice1 := make([]type, len)，这里 len 是数组的长度并且也是 slice 的初始长度

func main() {
	var s []int = make([]int, 5)
	// s[5] =100; error
	fmt.Printf("s: %v\n", s)
	fmt.Printf("s: %T\n", s)
	fmt.Printf("cap(s): %v\n", cap(s))
	/*
		s: [0 0 0 0 0]
		s: []int
		cap(s): 5
	*/

	// 因为字符串是纯粹不可变的字节数组，它们也可以被切分成 切片。
	str := "abcd"
	s1 := str[:]
	fmt.Printf("s1: %#v\n", s1)
	fmt.Printf("s1: %T\n", s1)
	// s1: "abcd"
	// s1: string
}
