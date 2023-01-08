package main

import "fmt"

// 定义一个 Rectangle 结构体，它的长和宽是 int 类型，
// 并定义方法 Area() 和 Primeter()，然后进行测试
type Rectangle struct{w,h int}

func Area(r *Rectangle)(area int){
	area = r.h * r.w
	return
}

func main(){
	r := new(Rectangle)
	r.h = 100
	r.w = 100

	fmt.Printf("Area(r): %v\n", Area(r))
}