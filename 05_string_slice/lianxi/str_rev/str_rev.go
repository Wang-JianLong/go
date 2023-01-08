package main

import (
	"fmt"
)

func main() {
	s := RevStr("Google")
	fmt.Println(s)

	s = RevStrBySl("你好呀")
	fmt.Printf("s: %v\n", s)
}

/*
编写一个程序，要求能够反转字符串，即将 “Google” 转换成 “elgooG”（提示：使用 []byte 类型的切片）。

如果您使用两个切片来实现反转，请再尝试使用一个切片（提示：使用交换法）。

如果您想要反转 Unicode 编码的字符串，请使用 []int32 类型的切片。
*/

func RevStr(str string) (s string) {
	if len(s) == 1 {
		return str
	}

	bs := []int32(str)
	for i := len(bs) - 1; i >= 0; i-- {
		s += string(bs[i])
	}
	return
}

func RevStrBySl(str string) string {
	// 转换为int切片
	is := []int32(str)

	for i := 0; i < len(is)/2; i++ {
		// is[i],is[len(is) - 1 -i] = is[len(is) - 1 -i],is[i]
		// 可以拆解为
		is[i] = is[len(is)-1-i] // 正着数 = 倒数
		tmp := is[len(is)-1 -i]
		is[len(is)-1-i] = tmp // 倒数 = 整数
	}
	return string(is)
}
