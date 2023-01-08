package main

import "fmt"

/*
结构体可以包含一个或多个 匿名（或内嵌）字段，
即这些字段没有显式的名字，只有字段的类型是必须的，
此时类型就是字段的名字。匿名字段本身可以是一个结构体类型，
即 结构体可以包含内嵌结构体。

可以粗略地将这个和面向对象语言中的继承概念相比较，
随后将会看到它被用来模拟类似继承的行为。
Go 语言中的继承是通过内嵌或组合来实现的，
所以可以说，在 Go 语言中，相比较于继承，组合更受青睐
*/

type innerS struct {
	in1, in2 int
}
type outerS struct {
	b int
	c float64
	int	// 同下
	innerS // 匿名字段
}

func main() {
	outers := &outerS{1,1,1,innerS{1,1}}
	outers2 := outerS{1,1,1,innerS{1,1}}

	fmt.Printf("outers: %v\n", outers)
	fmt.Printf("outers2: %v\n", outers2)
	/*
	outers: &{1 1 1 {1 1}}
	outers2: {1 1 1 {1 1}}
	*/
	fmt.Printf("outers.int: %v\n", outers.int)
	fmt.Printf("outers.innerS: %v\n", outers.innerS)

	// 可以直接写 匿名字段结构体的字段
	fmt.Printf("outers.in1: %v\n", outers.in1)
	fmt.Printf("outers.in2: %v\n", outers.in2)
	/*
	outers.int: 1
	outers.innerS: {1 1}

	通过类型 outer.int 的名字来获取存储在匿名字段中的数据，
	于是可以得出一个结论：
	在一个结构体中对于每一种数据类型
	只能有一个匿名字段
	*/
}
