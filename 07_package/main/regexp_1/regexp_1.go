package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pat := "[0-9]+.[0-9]+" //正则
	/*
		isPrent, err := regexp.Match("h", []byte(str))
		if err == nil {
			fmt.Printf("isPrent: %v\n", isPrent)
		}

		isPrent, err = regexp.MatchString("H", str)
		if err == nil {
			fmt.Printf("isPrent: %v\n", isPrent)
		}*/

	re, _ := regexp.Compile(pat)
	str1 := re.ReplaceAllString(str,"##.#")

	fmt.Printf("str1: %v\n", str1)// pat := "[0-9]+.[0-9]+" //正则

	// 参数是函数时
	str1 = re.ReplaceAllStringFunc(str,func(s string) string {
		return "*.*";
	})

	fmt.Printf("str1: %v\n", str1)//pat := "[0-9]+.[0-9]+" //正则
}
