package main

import "fmt"

/*
Go 标准库中许多地方都用了这个机制，
例如，json 包中的解码和 regexp 包中的 Complie 函数。
Go 库的原则是即使在包的内部使用了 panic，
在它的对外接口（API）中也必须用 recover 处理成返回显式的错误。
*/
func main() {
	println("main start")
	//	test()
	test2()
	println("main end")
}

func g() {
	panic("我的异常")
}
func test() {

	println("start.....")
	defer func() {
		println("done..........")
		if e := recover(); e != nil {
			println("异常信息 %v,%T", e, e)
		}
	}()
	g()
	println("end.....")
}

func test2() {
	fmt.Printf("\"test2 starting\": %v\n", "test2 starting")
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("我所要捕捉的异常%v\n", err)
		}
	}()
	g()
	fmt.Printf("\"test2 end...\": %v\n", "test2 end...")
}
