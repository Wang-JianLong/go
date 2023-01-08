package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// 设置一个结构体
type User struct {
	name string "username"
	id   int    "user id"
}

func (u *User) String() string {
	return strconv.Itoa(u.id) + "   " + u.name
}

var stru interface{} = User{"张三", 1}

func main() {
	v := reflect.ValueOf(stru)
	typ := reflect.TypeOf(stru)

	fmt.Printf("typ: %v\n", typ)
	knd := v.Kind() // knd: struct
	fmt.Printf("knd: %v\n", knd)
	for i := 0; i < v.NumField(); i++ {
		// 注意 value 类型与 type 类型不同
		fmt.Printf("i: %v  ifield: %v", i, v.Field(i)) // i: 0  ifield: 张三i: 1  ifield: 1
		f := typ.Field(i)
		fmt.Printf("name: %v  type:%v", f.Name, f.Type)
		// i: 0  ifield: 张三name: name  type:stringi: 1  ifield: 1name: id  type:interface
	}

	// 获取结构体方法
	//m := v.Method(0) 这里会报错 因为String方法写的是针对于指针的 而这里使用的是字面量
	//resu := m.Call(nil)

	resu := reflect.ValueOf(&User{"aa", 1}).Method(0).Call(nil)
	fmt.Printf("resu: %v\n", resu) // resu: [1   aa]
}
