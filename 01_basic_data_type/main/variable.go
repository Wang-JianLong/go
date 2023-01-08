package main

import (
	"fmt"
	"os"
	"runtime"
)

// 变量

// 声明全局变量
var (
	// 当一个变量被声明之后，系统自动赋予它该类型的零值：int 为 0，float 为 0.0，bool 为 false，string 为空字符串，指针为 nil。
	// 一个变量（常量、类型或函数）在程序中都有一定的作用范围，称之为作用域。如果一个变量在函数体外声明，则被认为是全局变量，可以在整个包甚至外部包（被导出后）使用
	a int
	b bool
	str string
)

func main(){
	// 声明变量的一般形式是使用 var 关键字：var identifier type。
	var name string = "zhangsan"
	// 尽管变量的标识符必须是唯一的，但你可以在某个代码块的内层代码块中使用相同名称的变量，则此时外部的同名变量将会暂时隐藏（结束内部代码块的执行后隐藏的外部同名变量又会出现，而内部同名变量则被释放），你任何的操作都只会影响内部代码块的局部变量
	if true {
		var name ="李四"
		fmt.Println(name) // 李四
	}
	fmt.Println(name) // 张三

	// 局部变量简短声明
	age := 19
	fmt.Println(age)

	oos()
}

func oos(){
	fmt.Println(runtime.GOOS) // windows
	fmt.Println(os.Getenv("GOROOT")) // C:\Program Files\Go
}