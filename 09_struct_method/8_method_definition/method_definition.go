package main

import "fmt"

/*
在 Go 语言中，结构体就像是类的一种简化形式，那么面向对象程序员可能会问：
	类的方法在哪里呢？在 Go 中有一个概念，它和方法有着同样的名字，
	并且大体上意思相同：Go 方法是作用在接收者（receiver）上的一个函数，
	`接收者是某种类型的变量。因此方法是一种特殊类型的函数`。

接收者类型可以是（几乎）任何类型，
	不仅仅是结构体类型：任何类型都可以有方法，
	甚至可以是函数类型，可以是 int、bool、string 或数组的别名类型。
	但是接收者不能是一个接口类型，
	因为接口是一个抽象定义，但是方法却是具体实现；
	如果这样做会引发一个编译错误：invalid receiver type…。

最后接收者不能是一个指针类型，但是它可以是任何其他允许类型的指针。

一个类型加上它的方法等价于面向对象中的一个类。
	一个重要的区别是：在 Go 中，类型的代码和绑定在它上面的方法的代码可以不放置在一起，
	它们可以存在在不同的源文件，唯一的要求是：它们必须是同一个包的。

类型 T（或 *T）上的所有方法的集合叫做类型 T（或 *T）的方法集。

因为方法是函数，所以同样的，
	不允许方法重载，即对于一个类型只能有一个给定名称的方法。
	但是如果基于接收者类型，是有重载的：
	具有同样名字的方法可以在 2 个或多个不同的接收者类型上存在，
	比如在同一个包里这么做是允许的：

*/

// 方法格式
// func (recv receiver_type) methodName(parameter_list) (return_value_list) { ... }
type Matrix struct{}
type denseMatrix struct{}
type sparseMatrix struct{}

// func (绑定的对象) name(...) [returnType] {}
func (a *denseMatrix) Add(b Matrix) Matrix  { return Matrix{} }
func (a *sparseMatrix) Add(b Matrix) Matrix { return Matrix{} }
func (m *Matrix) My() {
	// m == this ?
	fmt.Printf("m: %v\n", m)
}

func main() {
	m := Matrix{}
	d := &denseMatrix{}
	d.Add(m)

	s := &sparseMatrix{}
	s.Add(m)

	s1 := sparseMatrix{}
	s1.Add(m)

	fmt.Printf("\"end...\": %v\n", "end...")
	m.My()

	var st S = "AAAA"
	st.Get()
	// 理解下面
	S.SetS(st) // 类型名直接使用方法需要传递一个实例
	st.SetS() // 实例使用不需要
	// 这个方法的 receiver_type 并不是指针表示的

	sp := &st;
	sp.SetS() // AAAA 自动转化为字面量然后拷贝给方法
}

// 指针方法和值方法都可以在指针或非指针上被调用
type S string
// 为类型别名(接收)构建方法 
func (a *S) Get(){
	fmt.Printf("a: %v\n", a)
}
// 理解为静态方法 但是是值拷贝的
func (a S) SetS(){
	fmt.Printf("a: %v\n", a)
}
/*
如果 recv 是 receiver 的实例，Method1 是它的方法名，
	那么方法调用遵循传统的 object.name 选择器符号：recv.Method1()。

如果 recv 一个指针，Go 会自动解引用。
*/