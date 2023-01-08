package main

import (
	"fmt"
)

// 可变参数
/*
如果函数的`最后一个参数`是采用 ...type 的形式，那么这个函数就可以处理一个变长的参数，这个长度可以为 0，这样的函数称为变长函数

这个函数接受一个类似某个类型的 slice 的参数

*/
func main(){
	Count(1,23,3)
}

func Count(a ...int){
	count := 0 
	for i:=0; i < len(a);i++{
		count += a[i]
	}
	fmt.Println(count)
	// 变长参数可以作为对应类型的 slice 进行二次传递。
	printNums(a)
	// 使用结构
	test(1,2,A{p1:"1",p2:2})
	// 使用空接口
	test2([]interface{}{1,23,3})
}

func printNums(as []int){

}

func test(a int,b int , c A){

}
type A struct {
	p1 string
	p2 int
}

// 果一个变长参数的类型没有被指定，则可以使用默认的空接口 interface{}，这样就可以接受任何类型的参数

func test2(a interface{}){
	fmt.Println(a)
}