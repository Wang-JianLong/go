package main

import "fmt"

/*
定义一个结构体类型 Base，
	它包含一个字段 id，方法 Id() 返回 id，方法 SetId() 修改 id。

结构体类型 Person
	包含 Base，及 FirstName 和 LastName 字段。

结构体类型 Employee 包含一个 Person 和 salary 字段。

创建一个 employee 实例，然后显示它的 id
*/

type Base struct {
	id int
}

func (this *Base) Id() int    { return this.id }
func (this *Base) Set(id int) { this.id = id }

type Person struct{
	Base
	firstName,lastName string
}

type Employee struct {
	Person
	salary float64
}


func main() {
	e := &Employee{Person{Base{1},"wang","wu"},10000}
	fmt.Printf("e.Id(): %v\n", e.Id()) // e.Id(): 1
}
