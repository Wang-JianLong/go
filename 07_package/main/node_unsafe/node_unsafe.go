package main

import (
	"container/list"
	"fmt"
	"unsafe"
)

func main() {
	list := list.New()
	list.PushBack(101)
	list.PushBack(102)
	list.PushBack(103)
	for i := list.Front(); i!=nil;i = i.Next(){
		fmt.Printf("i: %v\n", i.Value)
	}
	var a int = 100;
	i := unsafe.Sizeof(a)
	fmt.Printf("i: %v\n", i)
}