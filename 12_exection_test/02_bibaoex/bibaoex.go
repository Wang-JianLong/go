package main

import "fmt"

func main() {
	println("main starting")
	f1 := errHandler(func(a, b int) {
		a = a / b
		println("f1 正常zhixing")
	})

	f2 := errHandler(func(a, b int) {
		panic("f2 exe....")
	})

	f1(1, 1)
	f2(1, 2)
	println("main end...")
}

/*
通过这种机制，所有的错误都会被 recover，并且调用函数后的错误检查代码也被简化为调用 check (err) 即可。在这种模式下，不同的错误处理必须对应不同的函数类型；它们（错误处理）可能被隐藏在错误处理包内部。可选的更加通用的方式是用一个空接口类型的切片作为参数和返回值。
*/
// 当大批量API是这种固定格式时，可以大批量调用
func errHandler(f func(a, b int)) func(int, int) {
	return func(i1, i2 int) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("错误是 %v\n", err)
			}
		}()
		f(i1, i2)

	}
}

// 这种方式可以 简化异常处理
func check(e error) {
	if e != nil {
		panic(e)
	}
}
