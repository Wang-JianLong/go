package main

// 应用闭包 将函数作为返回值

func main() {
	p := Add2()   // 返回一个函数赋值于p
	println(p(2)) // 3

	p2 := Add3(2)
	println(p2(2)) // 4

	p3 := Add4(2)
	println(p3(2)) // 4

	p4 := Add5(2)
	println(p4(2)) // 4
	println(p4(2)) // 6
	println(p4(2)) // 8
	println("================================================================")
	//
	f := fibonacci()
	for i := 0;i < 12;i++{
		println(f())
	}
}

// 返回一个 func(b int) int
func Add2() func(b int) int {
	return func(a int) int {
		return a + 1
	}
}

// 将外层函数的参数参与此函数内的值计算
func Add3(i int) func(i int) int {
	return func(a int) int {
		return i + a
	}
}

// Add3 变形
func Add4(i int) func(i int) int {
	x := i
	return func(j int) int {
		return x + j
	}
}

// 值保留
/*
闭包函数保存并积累其中的变量的值，不管外部函数退出与否，它都能够继续操作外部函数中的局部变量
*/
func Add5(i int) func(i int) int {
	x := i // 定义局部变量可以 // 得是在外部函数定义
	return func(j int) int {
		//i += j 函数参数也可以
		x += j
		return x
	}
}

func fibonacci() func() int {
	first:=0
	second:=1
	return func() int{
		defer func(){first,second=second,(first+second)}()
		return first//每次都返回第一个值，返回后再用defer进行计算已使数列向前推进
	}
}

