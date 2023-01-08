package main

import (
	"fmt"
	"reflect"
)

/*
结构体中的字段除了有名字和类型外，还可以有一个可选的标签（tag）：
	它是一个附属于字段的字符串，可以是文档或其他的重要标记。
	标签的内容不可以在一般的编程中使用，只有包 reflect 能获取它。
	它可以在运行时自省类型、属性和方法，
	比如：在一个变量上调用 reflect.TypeOf()
	可以获取变量的正确类型，如果变量是一个结构体类型，
	就可以通过 Field 来索引结构体的字段，然后就可以使用 Tag 属性。

*/

type User struct{
	id int "user id"
	name string "username"
}

func main(){
	u := User{1,"张三"}
	for i := 0; i < 2; i++ {
		utype := reflect.TypeOf(u)
		field := utype.Field(i)
		fmt.Printf("field.Tag: %v\n", field.Tag)
	}
}