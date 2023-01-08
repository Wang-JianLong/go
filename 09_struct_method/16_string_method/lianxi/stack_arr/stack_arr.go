package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
它的格子包含数据，比如整数 i、j、k 和 l 等等，格子从底部（索引 0）至顶部（索引 n）来索引。这个例子中假定 n=3，那么一共有 4 个格子。
一个新栈中所有格子的值都是 0。
push 将一个新值放到栈的最顶部一个非空（非零）的格子中。
pop 获取栈的最顶部一个非空（非零）的格子的值。现在可以理解为什么栈是一个后进先出（LIFO）的结构了吧
*/

// 为栈定义一 Stack 类型，并为它定义一个 Push 和 Pop 方法，再为它定义 String() 方法（用于调试）
// 它输出栈的内容，比如：[0:i] [1:j] [2:k] [3:l]
// 使用长度为 4 的 int 数组作为底层数据结构。

type Stack struct {
	index int
	datas [4]int
}

func (s *Stack) Push(data int) bool {
	if s.index <= len(s.datas) - 1 {
		// 没有满
		s.datas[s.index] = data
		s.index++
		return true
	}
	return false
}

func (s *Stack) Pop() (res int, isSuccess bool) {
	s.index--
	if s.index < 0 {
		isSuccess = false
		res = -1
		s.index++
		return
	}
	res = s.datas[s.index]
	isSuccess = true
	s.datas[s.index] = 0
	return
}

func (s *Stack) String() string {
	var sb strings.Builder
	for i, v := range s.datas {
		sb.WriteString("[" + strconv.Itoa(i) + ":" + strconv.Itoa(v) + "]")
	}
	return sb.String()
}

func main() {
	s := &Stack{index: 0}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	fmt.Printf("s: %v\n", s)
	s.Push(5)
	fmt.Printf("s: %v\n", s)

	for i := 0; i < 5; i++ {
		v, is := s.Pop()
		fmt.Printf("v:%v  is:%v\n", v, is)
		fmt.Printf("s: %v\n", s)
	}
	s.Push(4)
	fmt.Printf("s: %v\n", s)
}
