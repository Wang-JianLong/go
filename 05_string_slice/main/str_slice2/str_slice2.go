package main

import (
	"fmt"
	"unicode/utf8"
)

/*
Unicode 字符会占用 2 个字节，有些甚至需要 3 个或者 4 个字节来进行表示。
如果发现错误的 UTF8 字符，则该字符会被设置为 U+FFFD 并且索引向前移动一个字节。
和字符串转换一样，您同样可以使用 c := []int32(s) 语法，
这样切片中的每个 int 都会包含对应的 Unicode 代码，因为字符串中的每次字符都会对应一个整数。
	类似的，您也可以将字符串转换为元素类型为 rune 的切片：r := []rune(s)。
*/
func main(){
	s:= "abcdefg"
	i := []int32(s)
	fmt.Printf("i: %v\n", i)

	r := []rune(s)
	fmt.Printf("r: %v\n", r)
	// 可以通过代码 len([]int32(s)) 来获得字符串中字符的数量，
	// 但使用 utf8.RuneCountInString(s) 效率会更高一点
	count := utf8.RuneCountInString(s)
	fmt.Printf("count: %v\n", count)
	// 可以将字符串追加到字节切片的尾部

	b := []byte{1,2,3}
	b = append(b,s...)
	fmt.Printf("b: %v\n", b)

}