package main

import (
	"fmt"
	"io"
	"log"
)

// defer 修饰的func的参数 如果是函数会先执行该函数
func main() {
	fn()
	/*
	fn  入栈...
	fn......
	fn  出栈.*/

	func1("a")
	fn()
}

func fn() {
	name := "fn"
	defer un(trace(name))
	fmt.Println("fn......")
}

func trace(s string) (a string) {
	fmt.Println(s, " 入栈...")
	a = s
	return
}

func un(s string) {
	fmt.Println(s, " 出栈..")
}

func func1(s string)(n int,err error){
	defer func ()  {
		log.Printf("func1(%q) = %d,%v",s,n,err)
	}()

	return 7,io.EOF
}

//
type User struct{
	name string
}
