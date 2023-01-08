package main

import "fmt"

// Go 中的类型转换遵循严格的规则。当为结构体定义了一个 alias 类型时，
// 此结构体类型和它的 alias 类型都有相同的底层类型
type User struct {name string}
type MyUser User
func main(){
	mu := MyUser{"lisi"}	
	fmt.Printf("mu: %v\n", mu)

	fmt.Printf("User(mu): %v\n", User(mu))
	// mu: {lisi}
	// User(mu): {lisi}
}