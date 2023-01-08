package main

import "fmt"

func main() {
	str := "Hello World"
	p, s := StrSplit(str, 5)
	fmt.Printf("p: %v\n", p)
	fmt.Printf("s: %v\n", s)

	p1,s1 := StrSplitByByte(str,5)
	fmt.Printf("p1: %v\n", p1)
	fmt.Printf("s1: %v\n", s1)
}

/*
编写一个函数，要求其接受两个参数，原始字符串 str 和分割索引 i，然后返回两个分割后的字符串
*/
func StrSplit(str string, i int) (p string, s string) {
	sl := str[:]
	if len(sl) > i {
		p = sl[:i]
		s = sl[i:]
		return	
	}
	p = str
	s = ""
	return 
}
func StrSplitByByte(str string,i int)(p string, s string) {
	bs := []byte(str)
	if len(bs) > i {
		p = string(bs[:i])
        s = string(bs[i:])
        return
	}
	p = str
	s = ""
	return 
}

