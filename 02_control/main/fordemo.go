package main

import "fmt"

func main() {
	// 可以使用多个计数器 逆天
	N := 10
	for i, j := 0, N; i < j; i, j = i+1, j-1 {
		println(i, "===", j)
	}

	// 循环 Unicode 字符串
	str := "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n", len(str))
	for ix := 0; ix < len(str); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str[ix])
	}
	str2 := "日本語"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for ix := 0; ix < len(str2); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str2[ix])
	}

	// PrintGPlus()

	println("================================================================================================================================")
	str1 := "Go is a beautiful language!"
	for val, char := range str1 {
		println(val, "====", char)
	}
}

func PrintGPlus() {
	for i := 0; i < 25; i++ {
		for j := 1; j < i; j++ {
			print("G")
		}
		println()
	}
}
