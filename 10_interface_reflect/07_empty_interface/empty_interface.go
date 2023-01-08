package main

import "fmt"

var i = 5
var str = "ABC"

type Person struct {
	name string
	age  int
}

// 空接口或者最小接口 不包含任何方法，它对实现不做任何要求：
/*
任何其他类型都实现了空接口（它不仅仅像 Java/C# 中 Object 引用类型），any 或 Any 是空接口一个很好的别名或缩写。

空接口类似 Java/C# 中所有类的基类： Object 类，二者的目标也很相近。

可以给一个空接口类型的变量 var val interface {} 赋任何类型的值。

每个 interface {} 变量在内存中占据两个字长：
	一个用来存储它包含的类型，
	另一个用来存储它包含的数据或者指向数据的指针
*/
type Any interface{}

func main() {
	var val Any
	val = 5
	fmt.Printf("val has the value: %v\n", val)
	val = str
	fmt.Printf("val has the value: %v\n", val)
	pers1 := new(Person)
	pers1.name = "Rob Pike"
	pers1.age = 55
	val = pers1
	fmt.Printf("val has the value: %v\n", val)
	switch t := val.(type) {
	case int:
		fmt.Printf("Type int %T\n", t)
	case string:
		fmt.Printf("Type string %T\n", t)
	case bool:
		fmt.Printf("Type boolean %T\n", t)
	case *Person:
		fmt.Printf("Type pointer to Person %T\n", t)
	default:
		fmt.Printf("Unexpected type %T", t)
	}
}
