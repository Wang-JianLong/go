package main

import "fmt"

func main() {
	str := "1234"

	str = str[len(str)/2:] + str[:len(str)/2]

	fmt.Printf("str: %v\n", str)
}

/*
假设有字符串 str，那么 str[len(str)/2:] + str[:len(str)/2] 的结果是什么
*/