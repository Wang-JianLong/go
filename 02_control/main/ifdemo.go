package main

import "strconv"

func main() {
	if true {
		println("hello")
	} else {
		println("world")
	}
	username := "zhangsan"
	// if 的condtional 中可以定义变量 但是之间要用分号隔开
	if name := "zhangsan"; name == username {
		println("同名")
	} else {
		name = "lisi"
	}

	// println(name) 只能在if else 的作用域内

	i := "100-"
	num,error := strconv.Atoi(i)
	if error == nil{
		print(num)
	}else{
		print(error.Error())
	}
}
