package main

import "fmt"

/*
从字符串生成字节切片

假设 s 是一个字符串（本质上是一个字节数组），那么就可以直接通过 c := []byte(s) 来获取一个字节数组的切片 c。
另外，您还可以通过 copy 函数来达到相同的目的：copy(dst []byte, src string)

同样的，还可以使用 for-range 来获得每个元素
*/
func main() {
	s := "abcdefg"
	// 1 强转为字节数组获得字符串的切片
	s1 := []byte(s)
	fmt.Printf("s1: %v\n", s1)
	// 2 copy函数  // 注意dst 必须要初始化
	var s2 []byte = make([]byte, 7)
	copy(s2, s)
	fmt.Printf("s2: %v\n", s2)
	// 3 for range
	var s3 []byte
	for _,v := range s {
		s3 = append(s3, byte(v))
	}
	fmt.Printf("s3: %v\n", s3)
}
