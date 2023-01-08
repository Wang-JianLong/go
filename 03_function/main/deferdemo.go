package main

import "fmt"

// defer 和追踪

/*
关键字 defer 允许我们推迟到函数返回之前（或任意位置执行 return 语句之后）一刻才执行某个语句或函数（为什么要在返回之后才执行这些语句？因为 return 语句同样可以包含一些操作，而不是单纯地返回某个值）。

关键字 defer 的用法类似于面向对象编程语言 Java 和 C# 的 finally 语句块，它一般用于释放某些已分配的资源
*/
func main(){
	// function1()
	// a()
	// f()
	// 关键字 defer 允许我们进行一些函数执行完成后的收尾工作
	// defer file.Colse()

	// 加锁解锁
	/*
		mu.Lock()  
		defer mu.Unlock() 
	*/

	// b()
}

// 当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）
func f() {
	for i := 0; i < 5; i++ {
			defer fmt.Printf("%d ", i)
	}
}


// 使用 defer 的语句同样可以接受参数，下面这个例子就会在执行 defer 语句时打印 0 
func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}


func function1() {
	fmt.Printf("In function1 at the top\n")
	defer function2()
	fmt.Printf("In function1 at the bottom!\n")
}

func function2() {
	fmt.Printf("function2: Deferred until the end of the calling function!")
}

// 1个基础但十分实用的实现代码执行追踪的方案就是在进入和离开某个函数打印相关的消息
func trace(s string) { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }

func d() {
	trace("d")
	defer untrace("d")
	fmt.Println("in d")
}

func b() {
	trace("b")
	defer untrace("b")
	fmt.Println("in b")
	d()
}


