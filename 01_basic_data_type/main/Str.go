package main

import (
	"fmt"
	"unicode/utf8"
)

// 字符串
func main() {
	// 解释型字符串 会转义
	s1 := "a\r"
	fmt.Println(s1)
	// 非解释型字符串 不会转义
	s2 := `\r\n\t`
	fmt.Println(s2)
	// == 比较字符串是按照字节比较

	s3 := "56"
	s4 := "56"
	fmt.Println(s3 == s4) // true

	fmt.Println(len(s3))

	fmt.Println(s1[1]) // 按照索引获取字节

	fmt.Println(&s1)
	// 获取字符串中某个字节的地址的行为是非法的
	// fmt.Println(&s1[1])

	fmt.Println(CountStr("asSASA ddd dsjkdsjs dk"))
	fmt.Println(CountStr("asSASA ddd dsjkdsjsこん dk"))
}

/*
创建一个用于统计字节和字符（rune）的程序，并对字符串 asSASA ddd dsjkdsjs dk 进行分析，然后再分析 asSASA ddd dsjkdsjsこん dk，最后解释两者不同的原因（提示：使用 unicode/utf8 包
*/

func CountStr(str string) (rune, int) {
		return utf8.DecodeRuneInString(str)
}