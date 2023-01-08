package main

import "fmt"

// 定义一个结构体

// 类型 struct1 在定义它的包 pack1 中必须是唯一的，它的完全类型名是：pack1.struct1
type User struct {
	id int
	name string
}

// 或者
type UserBak struct {
	id,phone int
	_ string // 用不到的字段定义
}

// 使用new创建对象返回对象指针，而直接使用字面量创建对象返回对象本身 栈中和堆中
func main(){
	// 初始化方式

	// 1. 直接声明
	// 声明 var t User 也会给 t 分配内存，并零值化内存，但是这个时候 t 是类型
	var user User
	user.id = 1
	user.name = "张三"
	fmt.Printf("user: %v\n", user) // user: {1 张三}

	// 2. new 
	// 使用 new 函数给一个新的结构体变量分配内存，它返回指向已分配内存的指针：var t *T = new(T)，如果需要可以把这条语句放在不同的行

	// var userP *User = new(User)
	userP := new(User)
	// 变量 UserP 是一个指向 User 的指针，此时结构体字段的值是它们所属类型的零值。
	fmt.Printf("UserP: %v\n", userP) // UserP: &{0 }

	userP.id = 2
	userP.name = "李四"
	fmt.Printf("userP: %v\n", userP) //userP: &{2 李四}

	// 获取对象/实例的值
	fmt.Printf("user.id: %v\n", user.id)
	fmt.Printf("userP.id: %v\n", userP.id)

	// 3. 初始化一个结构体实例（一个结构体字面量：struct-literal）的更简短和惯用的方式如下
	// 混合字面量语法（composite literal syntax）&struct1{a, b, c} 是一种简写，
	// 底层仍然会调用 new ()，这里值的顺序必须按照字段顺序来写
	u := &User{10,"王五"}
	fmt.Printf("u: %v\n", u) //u: &{10 王五}
	fmt.Printf("u.id: %v\n", u.id)

	// 4. 或者
	s := User{11,"赵六"}
	fmt.Printf("s: %v\n", s) // s: {11 赵六}

	// 显式声明 顺序可以改变 某些字段可以省略
	d := User{id:1,name: "王二"}
	fmt.Printf("d: %v\n", d)
	// d: {1 王二}
	//g: &{0 李逵}
	g := &User{name:"李逵"}
	fmt.Printf("g: %v\n", g)
}

/*
注意：Go里的new和C++的new是不一样的：

    Go的new分配的内存可能在栈(stack)上，可能在堆(heap)上。
	C++ new分配的内存一定在堆上。Go的new分配的内存里的值是对应类型的零值，不能显示初始化指定要分配的值。
	C++ new分配内存的时候可以显示指定要存储的值。
	Go里没有构造函数，Go的new不会去调用构造函数。C++的new是会调用对应类型的构造函数。
*/
