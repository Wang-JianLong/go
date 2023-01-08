package main

import (
	"fmt"
	"strconv"
)

// 为 float64 定义一个别名类型 Celsius，
// 并给它定义 String()，它输出一个十进制数和 °C 表示的温度值
type Celsius float64

// TODO 字面值才能调用 为什么？
func (c Celsius) String() string{
	return strconv.FormatFloat(float64(c),'f',0,64) +"°C"
}

func main(){
	c  := Celsius(32.56)
	fmt.Printf("c: %v\n", c) // c: 33°C

	c2 := &Celsius2{33.5}
	fmt.Printf("c2: %v\n", c2) //c2: 34°C
}


type Celsius2 struct {
	float64
}

func (c *Celsius2) String() string{
	return strconv.FormatFloat(float64(c.float64),'f',0,64) +"°C"
}