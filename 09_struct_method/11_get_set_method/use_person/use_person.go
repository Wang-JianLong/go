package main

import (
	"09_struct_method/11_get_set_method/person"
	"fmt"
)

func main() {
	p := new(person.Person)
	// p.firstName = "A" Error 编译报错未定义 也就说 结构体导出 但是字段没有导出
	// 则需要 面向对象的GetSet
	p.SetFirstName("张三")
	fmt.Printf("p.FirstName(): %v\n", p.FirstName())
	fmt.Printf("p: %v\n", p)

}
