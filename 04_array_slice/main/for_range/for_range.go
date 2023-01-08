package main

import "fmt"

/*
第一个返回值 ix 是数组或者切片的索引，第二个是在该索引位置的值；他们都是仅在 for 循环内部可见的局部变量。value 只是 slice1 某个索引位置的值的一个拷贝，不能用来修改 slice1 该索引位置的值。*/
func main() {

}

func getArrayIndex() {
	ss := []string{"A", "B", "C"}
	// 如果只要索引 可以省略di
	for v := range ss {
		fmt.Printf("v: %v\n", v)
	}
}

// range 循环的value值无法改变被循环的原值
func TestRangevalue() {
	items := [...]int{10, 20, 30, 40, 50}
	for _, item := range items {
		item *= 2
	}
	fmt.Printf("items: %v\n", items)
}

// 通过使用省略号操作符 ... 来实现累加方法
func sum(items ...int) (result int) {
	for _, v := range items {
		result += v
	}

	return
}


