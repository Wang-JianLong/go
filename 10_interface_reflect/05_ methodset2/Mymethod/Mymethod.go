package main


// 定义两个接口

type Readr interface {
	Read()
}

type Writer interface {
	Write()
}

// 定义一个结构体
type User struct{}

// 此时为结构体实现这两个方法

// 接收者为 User结构体的字面量
func (s User) Read() {}

// 接收者为	User的指针
func (s *User) Write() {}

// TODO 接口实现的指针 与 结构体区别
func main() {
	// 测试
	// 作用于变量上的方法实际上是不区分变量到底是指针还是值的。
	// - 指 无论是指针还是字面量的 变量 都可以使用他们对应结构体的方法 无论接收者是指针还是字面量\

	// 当碰到接口类型值时，这会变得有点复杂，原因是接口变量中存储的具体值是不可寻址的

	// 1.定义一个存放字面量的变量
	u1 := User{}
	u1.Read()
	u1.Write()
	// 调用 结构体的方法都是可以的

	// OnWrite(u1) 报错
	OnWrite(&u1) // 得这样调用才行
	// cannot use u1 (variable of type User) as Writer value in argument to OnWrite: User does not implement Writer (method Write has pointer 
	// 接口方法定义在指针上
	OnRead(u1) // 结构体的字面量实现 Read接口

	// 2. 定义一个指针(引用)类型的变量
	u2 := &User{}
	
	OnWrite(u2) // User实现的Write方法的所有者是User指针 可以代表Writer
	
	OnRead(u2) // 虽然Read实现于User 字面量 但是 指针会被自动解引用。 所以 可以代表Readr

}

// 定义两个函数 分别接收两个接口
func OnRead(r Readr) {

}

func OnWrite(w Writer) {

}

/*
# 总结

在接口上调用方法时，必须有和方法定义时相同的接收者类型或者是可以从具体类型 

	- 指针方法可以通过指针调用
	- 值方法可以通过值调用
	- 接收者是值的方法可以通过指针调用，因为指针会首先被解引用
	- 接收者是指针的方法不可以通过值调用，因为存储在接口中的值(字面量)没有地址

将一个值赋值给一个接口时，编译器会确保所有可能的接口方法都可以在此值上被调用，
因此不正确的赋值在编译期就会失败。


译注

Go 语言规范定义了接口方法集的调用规则：

	- 类型 *T 的可调用方法集包含接受者为 *T 或 T 的所有方法集
	- 类型 T 的可调用方法集包含接受者为 T 的所有方法
	- 类型 T 的可调用方法集不包含接受者为 *T 的方法
*/
