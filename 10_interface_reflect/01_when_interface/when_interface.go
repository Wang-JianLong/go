package main

import "fmt"

/*
Go 语言不是一种 “传统” 的面向对象编程语言：它里面没有类和继承的概念。

但是 Go 语言里有非常灵活的 接口 概念，
	通过它可以实现很多面向对象的特性。接口提供了一种方式来 说明 对象的行为：
	如果谁能搞定这件事，它就可以用在这儿。

接口定义了一组方法（方法集），
	但是这些方法不包含（实现）代码：
	它们没有被实现（它们是抽象的）。接口里也不能包含变量。

接口定义格式
type Namer interface {
    Method1(param_list) return_type
    Method2(param_list) return_type
    ...
}
*/

// 定义一个接口
type Namer interface {
	M1()
	M2(a int)
	M3(a int) bool
}

/*
（按照约定，只包含一个方法的）接口的名字由方法名加 [e]r 后缀组成，
	例如 Printer、Reader、Writer、Logger、Converter 等等。
	还有一些不常用的方式（当后缀 er 不合适时），比如 Recoverable，
	此时接口名以 able 结尾，或者以 I 开头（像 .NET 或 Java 中那样）。

Go 语言中的接口都很简短，通常它们会包含 0 个、最多 3 个方法。
*/

func main1() {
	var ai Namer
	// 是一个多字（multiword）数据结构，它的值是 nil。它本质上是一个指针，虽然不完全是一回事。指向接口值的指针是非法的，它们不仅一点用也没有，还会导致代码错误
	// 编译通过 运行失败
	// ai.M1() // panic: runtime error: invalid memory address or nil pointer dereference
	fmt.Printf("namer: %v\n", ai) //namer: <nil>
}

/*
此处的方法指针表示通过`运行时反射能力`构建的。

类型（比如结构体）实现接口方法集中的方法，每一个方法的实现说明了`此方法是如何作用于该类型`的：即`实现接口`

同时方法集也构成了该类型的接口。

实现了 Namer 接口类型的变量可以赋值给 ai （接收者值），
	此时方法表中的指针会指向被实现的接口方法。

当然如果另一个类型（也实现了该接口）的变量被赋值给 ai，这二者（译者注：指针和方法实现）也会随之改变


- 类型不需要显式声明它实现了某个接口：接口被隐式地实现。多个类型可以实现同一个接口
- 实现某个接口的类型（除了实现接口方法外）可以有其他的方法。
- 一个类型可以实现多个接口。
- 接口类型可以包含一个`实例的引用`， 该实例的类型实现了此接口（接口是动态类型）。

即使接口在类型之后才定义，二者处于不同的包中，被单独编译：只要类型实现了接口中的方法，它就实现了此接口。
*/

// 例子		接口被隐式地实现
// 定义一个接口
type Eatr interface {
	Eat() string
	GetFoodName() string
}

// 定义一个结构体 拥有Eatr的方法 (或者说实现)
type Cat struct {
	name     string
	foodName string
}

func (self *Cat) Eat() string {
	return self.name + "喜欢吃" + self.foodName
}

func main2() {
	var eatr Eatr = &Cat{"田园猫", "汤泡饭和鱼"}

	fmt.Printf("eatr.Eat(): %v\n", eatr.Eat()) // eatr.Eat(): 田园猫喜欢吃汤泡饭和鱼
	fmt.Printf("eatr: value:%v 	type:%T", eatr, eatr)
	// eatr: value:&{田园猫 汤泡饭和鱼}	type:*main.Cat
}

/*
定义了一个结构体 Cat 和一个接口 Eatr，接口有一个方法 Eat() string。
定义了一个接收者类型是 Cat 方法的 Eat()
	- Cat 实现了 Eatr
所以可以将一个 Cat 类型的变量赋值给一个接口类型 Eatr的变量：var eatr Eatr = &Cat{"田园猫","汤泡饭和鱼"}

现在接口变量包含一个指向 Cat 变量的引用，通过它可以调用 Cat 上的方法 Eat()。
当然也可以直接在 Cat 的实例上调用此方法，
但是在接口实例上调用此方法更令人兴奋，它使此方法更具有一般性。
`接口变量里包含了接收者实例的值和指向对应方法表的指针`。


这是 多态 的 Go 版本，多态是面向对象编程中一个广为人知的概念：
	根据当前的类型选择正确的方法，
	或者说：同一种类型在不同的实例上似乎表现出不同的行为。
*/

// 此时 在接口上面多加一个方法 GetFoodName() string
// 如果此时 Cat没有 同样参数 同样返回值 并且指向Cat类型的方法时
// 编译器报错
// > cannot use &(Cat literal) (value of type *Cat) as Eatr value in variable declaration: *Cat does not implement Eatr (missing method GetFoodName)
// 那么就要Cat类型实现该方法
func (self *Cat) GetFoodName() string {
	return self.foodName
}

// 此时编译器不再报错

// 展示多态

// 定义一个接口 fly 拥有Fly方法
type Flyr interface {
	Fly()
}

// 定义一个函数 参数类型为接口  注意 函数参数是接口的话不能为定义为指针 【f.Fly undefined (type *Flyr is pointer to interface, not interface)】
func OnFly(f Flyr) {
	fmt.Printf("f: %T\n", f)
	f.Fly()
}

// 定义两个实现接口的结构体

type Sparrow struct{}

func (s *Sparrow) Fly() {
	fmt.Println("小麻雀在飞")
}

type Hawk struct{}

func (h *Hawk) Fly() {
	fmt.Println("老鹰在飞")
}

func main(){
	// 接口类型可以包含一个`实例的引用`， 该实例的类型实现了此接口（接口是动态类型）
	// var fly Flyr = Sparrow{} // cannot use (Sparrow literal) (value of type Sparrow) as Flyr value in variable declaration: Sparrow does not implement Flyr (method Fly has pointer receiver)
	// 不能在变量声明中使用（麻雀文字）（麻雀类型的值）作为传单值：麻雀不实现传单（方法飞有指针接收器）
	h := &Hawk{}
	s := &Sparrow{}
	flys := []Flyr{h,s}

	for _,f := range flys {
		OnFly(f)
	}
	/*
		f: *main.Hawk
		老鹰在飞
		f: *main.Sparrow
		小麻雀在飞
	*/
}