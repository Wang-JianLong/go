package main

import "fmt"

func main() {
	a := "a"
	var p *string = &a
	fmt.Println(*p) // a
	fmt.Println(p)  // 0xc000088230

	*p = "b"
	fmt.Println(a) // b

	fmt.Println(a == *(&a)) // true

	var c int = 0
	aa(c)
	println(c) // 0

	var d int = 0
	bb(&d)     // 这就是所谓的引用传递
	println(d) // 3
}

func aa(a int) {
	a = 3
}

func bb(a *int) {
	*a = 3
}
