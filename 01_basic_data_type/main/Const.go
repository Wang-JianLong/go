package main

import "fmt"
// 定义一个常量
// 因为在编译期间自定义函数均属于未知，因此无法用于常量的赋值，但内置函数可以使用，如：len ()。
const MAX_VALUE string = "MAX" // string可省略
const LEN = MAX_VALUE; //
// 数字型的常量是没有大小和符号的，并且可以使用任何精度而不会导致溢出
// 使用并行赋值的形式
const a,b,x = 1,2,3
// 常量还可以用作枚举：
const (
	One = 1
	Two = 2
)

// 自增值
const (
	c = iota + 50// 第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；
	// 简写为如下形式：
	d
	f // 2  // 52
	/*
	iota 也可以用在表达式中，如：iota + 50。在每遇到一个新的常量块或单个常量声明时， iota 都会重置为 0（ 简单地讲，每遇到一次 const 关键字，iota 就重置为 0 ）。
	*/
)

func main(){
	fmt.Println(f)
}