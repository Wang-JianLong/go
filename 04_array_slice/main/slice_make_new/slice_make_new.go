package main

import "fmt"

// 看起来二者没有什么区别，都在堆上分配内存，但是它们的行为不同，适用于不同的类型
func main() {
	// new (T) 为每个新的类型 T 分配一片内存，初始化为 0 并且返回类型为 * T 的内存地址：这种方法 返回一个指向类型为 T，值为 0 的地址的指针，它适用于值类型如数组和结构体
	// 它相当于 &T{}
	n := new([]int)
	fmt.Printf("n: %#v\n", n) // n: &[]int(nil)
	fmt.Printf("n: %T\n", n)  // n: *[]int
	// (*n)[1] = 100 eror
	np := *n
	// np[1] = 100
	fmt.Printf("np: %#v\n", np) // np: []int(nil)

	// make(T) 返回一个类型为 T 的初始值，它只适用于 3 种内建的引用类型：切片、map 和 channel

	// 换言之，new 函数分配内存，make 函数初始化

	s := make([]int, 10, 50)
	fmt.Printf("s: %#v\n", s)          //s: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	fmt.Printf("cap(s): %v\n", cap(s)) //cap(s): 50
	fmt.Printf("len(s): %v\n", len(s)) //len(s): 10

	s[9] = 100
	fmt.Printf("s: %v\n", s) //s: [0 0 0 0 0 0 0 0 0 100]

	fmt.Println(GetByteLen())

	//GetByteLen2()
}

// TODO 未解决
func GetByteLen() (lens int, caps int) {
	b := make([]byte, 5)
	fmt.Printf("b: %v\n", b)           // b: [0 0 0 0 0]
	fmt.Printf("len(b): %v\n", len(b)) //5
	fmt.Printf("cap(b): %v\n", cap(b)) //5
	b = b[2:4]
	fmt.Printf("b: %v\n", b)           // b: [0 0]
	fmt.Printf("len(b): %v\n", len(b)) //2
	fmt.Printf("cap(b): %v\n", cap(b)) //3
	lens = len(b)                      //2
	caps = cap(b)                      //3 ???
	return
}

func GetByteLen2() {
	s1 := []byte{'p', 'o', 'e', 'm'}
	s2 := s1[2:4] // s1[2:]
	fmt.Printf("s1: %#v\n", s1)
	fmt.Printf("s2: %#v\n", s2)
	s2[1] = 't'
	fmt.Printf("s1: %#v\n", s1)
	fmt.Printf("s2: %#v\n", s2)
}
