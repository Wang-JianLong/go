package main

import "fmt"
import "07_package/main/pack/pack2"

/*
	包的初始化顺序
	- 导入包
	- main包
	- main函数

	一个没有导入的包将通过分配初始值给所有的包级变量和调用源码中定义的包级 init 函数来初始化
	一个包可能有多个 init 函数，它们甚至可以存在于同一个源码文件中。它们的执行是无序的。这是测定包的值是否只依赖于相同包下的其他值或者函数的最好的例子

	init 函数是不能被调用的
	导入的包在包自身初始化前被初始化，而一个包在程序执行中只能初始化一次。

一个包下可以有多个源文件

*/
func main() {
	fmt.Printf("Main get pack2.Name: %v\n", pack2.Name)
	fmt.Printf("pack2.Go: %v\n", pack2.Go)
	
}

func init() {
	fmt.Printf("\"Main...init\": %v\n", "Main...init")
}
