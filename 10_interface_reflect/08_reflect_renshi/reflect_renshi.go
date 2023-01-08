package main

import (
	"fmt"
	"reflect"
)

/*
#可以通过反射来分析一个结构体。
#反射是用程序检查其所拥有的结构，尤其是类型的一种能力；
#这是元编程的一种形式。
#
#反射可以在运行时检查类型和变量，例如它的大小、方法和 动态的调用这些方法。
#这对于没有源代码的包尤其有用。这是一个强大的工具，除非真得有必要，否则应当避免使用或小心使用。
*/

func main() {
	var f float64 = 1.0

	// 实际上，反射是通过检查一个接口的值，变量首先被转换成空接口。
	v := reflect.ValueOf(f)
	t := reflect.TypeOf(f)
	fmt.Printf("v: %T %v\n", v, v) // v: reflect.Value 1
	fmt.Printf("t: %T %v\n", t, t) // t: *reflect.rtype float64

	fmt.Printf("t.Kind().String(): %v\n", t.Kind().String()) //t.Kind().String(): float64
	// 接口的值包含一个 type 和 value。
	// 反射可以从接口值反射到对象，也可以从对象反射回接口值。
	fmt.Printf("v.Float(): %v\n", v.Float()) // 1

	// Kind 总是返回底层类型
	// Type Interface() 方法可以得到还原（接口）值
	fmt.Println(v.Interface()) //1
	
	fmt.Printf("\"a\": %v\n", "a")
	fmt.Printf("\"b\": %v\n", "b")
	fmt.Printf("\"c\": %v\n", "c")

}
