package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 类型别名
type Rope string

func main() {
	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens) // 只是以纳秒作为一个种子
	for i := 0; i < 10; i++ {
		// rand.Float32() 得到 0.0~1.0 之间的随机数
		fmt.Printf("%2.2f / ", 100*rand.Float32())
	}
	var name Rope = "Rope"
	fmt.Println(name)

	var ch int = '\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00101234'
	fmt.Printf("%d - %d - %d\n", ch, ch2, ch3) // integer
	fmt.Printf("%c - %c - %c\n", ch, ch2, ch3) // character
	fmt.Printf("%X - %X - %X\n", ch, ch2, ch3) // UTF-8 bytes
	fmt.Printf("%U - %U - %U", ch, ch2, ch3)   // UTF-8 code point
}
