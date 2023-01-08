package main

// 闭包
func main() {
	// 1. 变量接收函数
	fn := func(x int, y int) int {
		return x + y
	}
	println(fn(1,2))

	// 2. 匿名函数
	v := func(x int, y int) int {return x * y}(1,3)
	println(v)
	// 3. 

	println(f())
}

func f() (ret int) {
	defer func() {
			ret++
	}()
	// 命名返回值时间上 发返回值时将 1 赋值给 ret 
	return 1
}