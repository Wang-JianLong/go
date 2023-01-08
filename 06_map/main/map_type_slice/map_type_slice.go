package main

import "fmt"

/*
假设我们想获取一个 map 类型的切片，
我们必须使用两次 make() 函数，第一次分配切片，第二次分配 切片中每个 map 元素
*/
func main(){
	// 至为map的切片
	a := make([]map[int]int,5)
	for i := range a{
		a[i] = make(map[int]int,1)
		a[i][i] = i
	}

	fmt.Printf("a: %v\n", a)
	// a: [map[0:0] map[1:1] map[2:2] map[3:3] map[4:4]]

	b := make([]map[int]int,5)

	for _,v := range b{
		v = make(map[int]int)
		v[1] = 2
	}
	// 需要注意的是，应当像 A 版本那样通过索引使用切片的 map 元素。
	// 在 B 版本中获得的项只是 map 值的一个拷贝而已，所以真正的 map 元素没有得到初始化。
	fmt.Printf("b: %v\n", b)
	// b: [map[] map[] map[] map[] map[]]

}