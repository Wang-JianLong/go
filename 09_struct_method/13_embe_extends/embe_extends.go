package main

import "fmt"

/*
当一个匿名类型被内嵌在结构体中时，匿名类型的可见方法也同样被内嵌，
这在效果上等同于外层类型 继承 了这些方法：将父类型放在子类型中来实现亚型。
这个机制提供了一种简单的方式来模拟经典面向对象语言中的子类和继承相关的效果，
也类似 Ruby 中的混入（mixin）


内嵌将一个已存在类型的字段和方法注入到了另一个类型里：
	匿名字段上的方法 “晋升” 成为了外层类型的方法。
	当然类型可以有只作用于本身实例而不作用于内嵌 “父” 类型上的方法，

可以覆写方法（像字段一样）：
	和内嵌类型方法具有同样名字的外层类型的方法会覆写内嵌类型对应的方法。

因为一个结构体可以嵌入多个匿名类型，所以实际上我们可以有一个简单版本的多重继承，
就像：type Child struct { Father; Mother}

结构体内嵌和自己在`同一个包中的结构体时，可以彼此访问对方所有的字段和方法`。
*/

type Engien interface {
	Start()
	Stop()
}

type Car struct {
	wheelCount int
	Engien
}

type Mercedes struct {
	Car
}

func (self *Car) GoToWorkIn() {
	self.Start()
	self.Stop()
}

func (self *Car) numberOfWheels() int {
	return self.wheelCount
}

func (self *Car) Start() {
	fmt.Printf("start...\n")
}

func (self *Car) Stop() {
	fmt.Printf("stop...\n")
}

func (self * Mercedes) sayHiToMerkel(){
	fmt.Printf("\"HiToMerkel\": %v\n", "HiToMerkel")
}

func main() {
	car := new(Car)
	car.GoToWorkIn()

	m := &Mercedes{*car}
	m.sayHiToMerkel()
	m.Stop()
}

/*
内嵌的类型不需要指针，Customer 也不需要 Add 方法，它使用 Log 的 Add 方法，Customer 有自己的 String 方法，并且在它里面调用了 Log 的 String 方法。

如果内嵌类型嵌入了其他类型，也是可以的，那些类型的方法可以直接在外层类型中使用。

因此一个好的策略是创建一些小的、可复用的类型作为一个工具箱，用于组成域类型
*/