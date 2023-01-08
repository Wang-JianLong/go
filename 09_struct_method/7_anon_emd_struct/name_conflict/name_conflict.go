package main

import "fmt"

/*
当两个字段拥有相同的名字（可能是继承来的名字）时该怎么办呢？

	外层名字会覆盖内层名字（但是两者的内存空间都保留），
	这提供了一种重载字段或方法的方式；

	如果相同的名字在同一级别出现了两次，如果这个名字被程序使用了，
	将会引发一个错误（不使用没关系）。
	没有办法来解决这种问题引起的二义性，必须由程序员自己修正
*/

type A struct {
	a,b int
}

type B struct {
	a,b int
	A
}

// 同一级别出现两次
type C struct {
	A
	B
}

type D struct {B; b float32}

func main() {
	b := &B{2,2,A{1,1}}
	// 外层包围内层
	fmt.Printf("b.a: %v\n", b.a) // 2
	fmt.Printf("b.b: %v\n", b.b) // 2

	// 构建正常
	c := &C{A{1,1},B{2,2,A{3,3}}}
	fmt.Printf("c: %v\n", c)  //c: &{{1 1} {2 2 {3 3}}}
	// 当访问数据时
	// c.a 编译报错 
	// 需要这样访问
	fmt.Printf("c.A.a: %v\n", c.A.a) // 1

	d := &D{B{2,2,A{}},1.0}
	// 这样访问是可以的 D的b是float类型 而B的b是int类型 类型不同可以
	fmt.Printf("d.b: %#v  %T\n", d.b,d.b) // d.b: 1  float32
}
