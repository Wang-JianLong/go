// panic_defer.go
package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f.")

}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)

		}

	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")

}

// 注意 defer 修饰的函数 是最后执行 的就是
// 即使所处函数 调用了另一个函数  后进先出原则的栈结构
func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))

	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)

}
