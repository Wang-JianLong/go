package main

import (
	"fmt"
	"strings"
)

type User struct{ firstName, lastName string }

func main() {

	// a := 1
	// p := &a
	// fmt.Printf("p: %v\n", p) // 0x....
	// (*p) = 4
	// fmt.Printf("a: %v\n", a) // 4

	user1 := User{"wang","wu"}
	upUserNoPointer(user1)
	fmt.Printf("user1: %v\n", user1)
	upUserOnPointer(&user1)
	fmt.Printf("user1: %v\n", user1)
	
	user2 := new(User)
	user2.firstName ,user2.lastName= "wang","wu"
	(*user2).firstName = "wang1" // 这是合法的 通过解指针的方式来设置值
	upUserOnPointer(user2)
	fmt.Printf("user2: %v\n", user2)

	user3 := &User{"wang","wu"}
	upUserOnPointer(user3)
	fmt.Printf("user3: %v\n", user3)
}

// 引用传递
func upUserOnPointer(user *User) {
	user.firstName = strings.ToUpper(user.firstName)
	user.lastName = strings.ToUpper(user.lastName)
}

// 值传递 无法改变结构体说指向的内存地址
func upUserNoPointer(user User) {
	user.firstName = strings.ToUpper(user.firstName)
	user.lastName = strings.ToUpper(user.lastName)
}

/*
Go 语言中，结构体和它所包含的数据在内存中是以连续块的形式存在的，
即使结构体中嵌套有其他的结构体，这在性能上带来了很大的优势。

不像 Java 中的引用类型，一个对象和它里面包含的对象可能会在不同的内存空间中，
这点和 Go 语言中的指针很像
*/