package main

import (
	. "09_struct_method/16_string_method/lianxi/stack_struct"
	"fmt"
)
func main(){
	s := NewStack()
	fmt.Println(s.Pop())

	for i := 1; i < 5; i++ {
		s.Push(i)
		fmt.Printf("s: %v\n", s)
	}

	for i := 0; i < 5; i++ {
		fmt.Println(s.Pop())
		fmt.Printf("s: %v\n", s)
	}

}