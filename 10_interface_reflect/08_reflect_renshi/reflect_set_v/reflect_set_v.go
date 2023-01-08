package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.3
	v := reflect.ValueOf(x)

	fmt.Printf("v.CanSet(): %v\n", v.CanSet())

	// 函数通过传递一个 x 拷贝创建了 v，那么 v 的改变并不能更改原始的 x
	// 函数通过传递一个 x 拷贝创建了 v，那么 v 的改变并不能更改原始的 x。要想 v 的更改能作用到 x，那就必须传递 x 的地址
	v = reflect.ValueOf(&x)
	if v.CanSet() {
		v.SetFloat(64)
		fmt.Printf("x: %v\n", x)
		// 但是这个是任然不行的
	}
	// 需要的到它的指针
	p := v.Elem()
	// Elem函数会间接的使用指针
	if p.CanSet() {
		p.SetFloat(100)
		fmt.Printf("x: %v\n", x) // 修改成功 x = 100
	}
	// 查看一下 p的类型
	fmt.Printf("p: %v\n", p)
	// 反射中有些内容是需要地址去改变它的状态的
}
