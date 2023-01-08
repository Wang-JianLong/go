package main

import (
	"fmt"
	"reflect"
)

type A struct {
	Name string
}

type B struct {
	name string
}

func main() {
	a := A{"张三"}
	b := B{"lisi"}

	v1 := reflect.ValueOf(a)
	v2 := reflect.ValueOf(b)

	b1 := v1.Field(0).CanSet()
	b2 := v2.Field(0).CanSet()

	fmt.Println(b1) // false
	fmt.Println(b2) // false

	// 注意使用地址改变值
	v3 := reflect.ValueOf(&a)
	v4 := reflect.ValueOf(&b)

	b3 := v3.Elem().Field(0).CanSet()
	b4 := v4.Elem().Field(0).CanSet()

	fmt.Printf("b3: %v\n", b3) // true
	fmt.Printf("b4: %v\n", b4) // false 注意 属性字段首字母大写才能Set值


	// 注意 没有Elem的话 结构体获取字段会报错
	b5 := reflect.ValueOf(&a).Field(0).CanSet()
	b6 := reflect.ValueOf(&b).Elem().Field(0).CanSet()
	fmt.Printf("b5: %v\n", b5)
	fmt.Printf("b6: %v\n", b6)

}
