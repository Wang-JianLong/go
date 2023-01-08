package main

import (
	"crypto/sha1"
	"fmt"
	"io"
)
// 通过调用 sha1.New() 创建了一个新的 hash.Hash 对象，用来计算 SHA1 校验值。Hash 类型实际上是一个接口，它实现了 io.Writer 接口

func main() {
	hasher := sha1.New()
	io.WriteString(hasher,"test")
	b := []byte{}
	sum := hasher.Sum(b)
	
	fmt.Printf("hasher.Sum(b): %x\n", sum)// 输出十六进制字符串
	fmt.Printf("hasher.Sum(b): %d\n", hasher.Sum(b))

	hasher.Reset() // 重置状态
	io.WriteString(hasher,"aaaaaa")

	fmt.Printf("hasher.Sum(b): %v\n", hasher.Sum(b))
}