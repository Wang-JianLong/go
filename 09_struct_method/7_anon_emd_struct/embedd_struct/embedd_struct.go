package main

import "fmt"

/*
同样地结构体也是一种数据类型，
所以它也可以作为一个匿名字段来使用，
如同上面例子中那样。
外层结构体通过 outer.in1 直接进入内层结构体的字段，
内嵌结构体甚至可以来自其他包。
内层结构体被简单的插入或者内嵌进外层结构体。
这个简单的 “继承” 机制提供了一种方式，
使得可以从另外一个或一些类型继承部分或全部实现。
*/
type A struct {
	a, b int
}

type B struct {
	d, c int
	A    // 也就是说 可以理解为继承A的全部字段
}

func main() {
	b := new(B)
	b.a, b.b, b.c, b.d = 1, 2, 3, 4
	fmt.Printf("b: %v\n", b)
	fmt.Printf("b.A: %v\n", b.A)
	/*
	b: &{4 3 {1 2}}
	b.A: {1 2}
	*/
}
