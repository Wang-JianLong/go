package main

import "fmt"

type M interface {
	m()
}

type X interface {
	x()
}

type Y struct {
}

func (s *Y) x() { fmt.Printf("X s: %T %v\n", s, s) }
func (s *Y) m() { fmt.Printf("M s: %T %v\n", s, s) }


// 只要 x 的底层类型（动态类型）定义了 print 方法这个调用就可以正常运行
func main() {
	var m1 M
	var x1 X

	y := new(Y) // 实现了M X

	var e interface{}

	e = y

	m1, _ = e.(M)
	m1.m()

	x1, _ = m1.(X)
	x1.x()
	e = x1
}
