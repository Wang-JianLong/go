package main

import "fmt"

func main() {
	test2()
}
func test(){
		// Go 语言中的数组是一种 值类型（不像 C/C++ 中是指向首元素的指针），所以可以通过 new() 来创建
		p := new([5]int)
		fmt.Printf("p: %#v\n", p) // p: &[5]int{0, 0, 0, 0, 0}
		fmt.Printf("p: %T\n", p) // p: *[5]int
	
		// 这样的结果就是当把一个数组赋值给另一个时，需要再做一次数组内存的拷贝操作
		arr := *p;
		arr[2] = 100;
		// 这样两个数组就有了不同的值，在赋值后修改 arr2 不会对 arr1 生效
		fmt.Printf("arr: %#v\n", arr) // arr: [5]int{0, 0, 100, 0, 0}
		fmt.Printf("p: %#v\n", p) // p: &[5]int{0, 0, 0, 0, 0}
}
// 如果想修改原数组，那么 arr2 必须通过 & 操作符以引用方式传过来
// 另一种方法就是生成数组切片并将其传递给函数
func test2(){
	var arr [3]int;
	f(arr)
	fmt.Printf("arr: %v\n", arr) // arr: [0 0 0]
	fp(&arr)
	fmt.Printf("arr: %v\n", arr) // arr: [0 0 300]
}

func f(arr [3]int){
	arr[2] = 100
}

func fp(arr *[3]int){
	arr[2] = 300
}