package main

import "fmt"

/*
为 int 定义别名类型 TZ，定义一些常量表示时区
，比如 UTC，定义一个 map，它将时区的缩写映射为它的全称，
比如：UTC -> "Universal Greenwich time"。
为类型 TZ 定义 String() 方法，它输出时区的全称。
*/
const (
	UTC    = 1
	Us  TZ = iota
	Uk  TZ = iota + 1
	Cn  TZ = iota + 2
	As  TZ = iota + 3
)

type TZ int

func (t TZ) String() string {
	allname := [4]string{"美国", "英国", "中国", "澳大利亚"}
	return allname[t]
}

func main() {
	var t TZ = Uk
	fmt.Printf("t: %v\n", t)

}
