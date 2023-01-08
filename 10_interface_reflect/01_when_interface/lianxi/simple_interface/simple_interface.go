package main

import "fmt"

/*

定义一个接口 Simpler，
	它有一个 Get() 方法和一个 Set()
	，Get() 返回一个整型值，Set() 有一个整型参数。
	创建一个结构体类型 Simple 实现这个接口。

接着定一个函数，
它有一个 Simpler 类型的参数，
调用参数的 Get() 和 Set() 方法。
在 main 函数里调用这个函数，看看它是否可以正确运行。
*/

type Simpler interface{
	Get() int
	Set(data int)
}

type Simple struct{
	id int
}

func (self *Simple) Get() int{
	return self.id
}

func (self *Simple) Set(id int){
	self.id = id
}

func On(s Simpler){
	s.Set(1)
	fmt.Printf("s.Get(): %v\n", s.Get())
	fmt.Printf("s: %v\n", s)
}

func main() {
	s := &Simple{}
	On(s)
	/*
		s.Get(): 1
		s: &{1}
	*/
}
