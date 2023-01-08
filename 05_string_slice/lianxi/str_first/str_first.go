package main

import "fmt"

// 编写一个程序，要求能够遍历一个数组的字符，并将当前字符和前一个字符不相同的字符拷贝至另一个数组
func main() {
	b := check([]byte("Googleasdlkjkk;aassdhsjjd"))
	fmt.Printf("string(b): %v\n", string(b))
}

func check(data []byte) []byte {
	// 先将第一个拷贝进去
	c := append(make([]byte,0),data[0])
	// 遍历出后一个与前一个比较是否不相同 是的话就append至新切片
	for i := 1; i < len(data); i++ {
		if data[i] != data[i-1] {
			c = append(c, data[i])
		}
	}
	return c
}
