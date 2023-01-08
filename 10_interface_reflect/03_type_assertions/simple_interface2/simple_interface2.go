package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type RSimple struct {
}

func (r *RSimple) Area() float32 {
	return 1.0
}

/*
接着练习 11.1 中的内容，创建第二个类型 RSimple，它也实现了接口 Simpler，
写一个函数 fi，使它可以区分 Simple 和 RSimple 类型的变量。
*/
func fi(s Shaper) {
	switch s.(type) {
	case *RSimple:
		fmt.Println("这是rs")
	case *Square:
		fmt.Println("这是Sq")
	}
}

func main() {
	s := &Square{}
	r := &RSimple{}

	fi(s)
	fi(r)
}
