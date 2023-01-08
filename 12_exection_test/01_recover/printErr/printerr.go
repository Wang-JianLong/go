package main

import "fmt"

func main() {
	e := fmt.Errorf("%v\n", "这是我自定义的错误")
	fmt.Printf("e: %v\n", e)
	fmt.Printf("e的类u: %T\n", e)

}
