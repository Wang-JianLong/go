package main

import "fmt"

func main() {
	s := "abcd"
	/*
		使用 substr := str[start:end]
		可以从字符串 str 获取到从索引 start 开始到 end-1 位置的子字符串。
		同样的，str[start:] 则表示获取从 start 开始到 len(str)-1 位置的子字符串。
		而 str[:end] 表示获取从 0 开始到 end-1 的子字符串。*/

	sub := s[:]
	fmt.Printf("sub: %v\n", sub)
	// 字符串是不可变的	但是可以转换字符切片更改

	bs := []byte(s)

	bs[1] = 'g'
	s = string(bs)
	fmt.Printf("s: %v\n", s)
}