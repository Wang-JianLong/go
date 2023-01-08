package main

import (
	"09_struct_method/5_struct_factory/facotry"
	"fmt"
)

func main() {
	// 访问失败
	// facotry.max
	m := facotry.NewMax()
	fmt.Printf("m: %v\n", m)
}
