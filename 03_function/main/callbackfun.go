package main

import "fmt"

// 函数可以作为其它函数的参数进行传递，然后在其它函数内调用执行，一般称之为回调
func main() {
	v:=callback(1,Add)
	fmt.Printf("%d",v)
}

func Add(a, b int)  (c int){
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
	c= a+b
	return 
}

func callback(y int, f func(int, int) int) int{
	return f(y, 2) // this becomes Add(1, 2)
}
