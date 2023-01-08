package main

import (
	"fmt"
)

/*
定义结构体 employee，它有一个 salary 字段，
给这个结构体定义一个方法 giveRaise 来按照指定的百分比增加薪水
*/

type Employee struct {
	salary float64
}

func (this *Employee) GiveRaise(f float64) {
	this.salary *= f
}

func main() {
	e := &Employee{10000}
	e.GiveRaise(10)
	fmt.Printf("e: %v\n", e)
}
