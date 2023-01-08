package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	// stringBuffer()
	// TwoSlice()
	bs := [20]byte{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}

	s1 := bs[:10]
	fmt.Printf("s1: %v\n", s1)
	s2 := bs[10:]
	fmt.Printf("s2: %v\n", s2)
	sl := Append(s1, s2)
	fmt.Printf("sl: %v\n", sl)
}

func DingYi() {
	// 1 定义 获得一个指针
	// 1.1 定义
	// var buf bytes.Buffer;
	// 1.2 new(bytes.Buffer)
	// var buf *bytes.Buffer = new(bytes.Buffer)
	// 1.3 函数获取 创建一个 Buffer 对象并且用 buf 初始化好；NewBuffer 最好用在从 buf 读取的时候使用
	// var buf = bytes.NewBuffer([]byte{'a','b','c'})
}

// 这种实现方式比使用 += 要更节省内存和 CPU，尤其是要串联的字符串数目特别多的时候
func stringBuffer() {
	var buf bytes.Buffer
	str := "abcdefg"
	for i, char := range str {
		buf.WriteString(strconv.Itoa(i) + "." + strconv.QuoteRune(char))
	}

	fmt.Printf("buf.String(): %v\n", buf.String()) // buf.String(): 0.'a'1.'b'2.'c'3.'d'4.'e'5.'f'6.'g'
}

// 把一个缓存 buf 分片成两个 切片：第一个是前 n 个 bytes，后一个是剩余的，用一行代码实现
func TwoSlice() {
	var buf bytes.Buffer
	buf.WriteString("abcdefghik")

	one, two := buf.Bytes()[:5], buf.Bytes()[5:]

	fmt.Printf("one: %v\n", one)
	fmt.Printf("two: %v\n", two)
}

// 给定切片 sl，将一个 []byte 数组追加到 sl 后面。写一个函数 Append(slice, data []byte) []byte，该函数在 sl 不能存储更多数据的时候自动扩容
func Append(slice, data []byte) []byte {
	var buf bytes.Buffer
	// 将 slice 写入 buf
	buf.Write(slice)
	// 将 []byte 写入 buf
	for _, v := range data {
		buf.WriteByte(v)
	}
	// 获取总长度
	num := len(slice) + len(data)
	sl := make([]byte, num)
	// 将 buf 数据写入新的slice
	buf.Read(sl)
	return sl
}
