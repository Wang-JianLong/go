package main

import (
	"bytes"
	"fmt"
)

// import wjl "hello_world/main/wjl"
// 注释
/*
	多行注释
*/
func main() {
	/* fmt.Println("Hello, world!")
	wjl.Geta()
	fmt.Println(getSum(1,3))
	a,b := getSums()
	fmt.Println(a,b) */
	var a IZ = 3
	fmt.Println(a)
	var b bytes.Buffer
	b.WriteString("a")
	fmt.Printf("b.String(): %v\n", b.String())
	fmt.Printf("\"Hello\": %v\n", "Hello")

	s := []int{1, 2, 3}
	s = append(s, 100)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	// 现在是幻想时间
	fmt.Printf("s: %v\n", s)
	GetErrorString()
}

type IZ int
d

func init() {
	fmt.Println("init....")
}
func getSum(a int16, b int16) int16 {
	return a + b
}

func getSums() (int, int) {
	return 1, 2
}

func Getwjl() string {
	return "wjl"

}
