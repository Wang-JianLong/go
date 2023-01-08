package main

import (
	"fmt"
	"wjl/inter"
)

type Array []int

func (a Array) Len() int {
	return len(a)
}
func (a Array) Less(i, j int) bool {
	return a[i] < a[j]
}
func (a Array) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	var a []int = []int{12, 23, 12, 1, 3, 54, 1, 34, 46, 323}
	fmt.Printf("a: %v\n", a)
	array := Array(a)
	inter.SortFun(array)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("GetOk(): %v\n", GetOk())

	i := &inter.InnerImpl{}
	i.Move()
}


func GetOk()string{
	return "ok"
}
