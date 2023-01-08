package main

import "fmt"

func main() {
	println("main start")
	divide := Divide(func(i1, i2 int) int {
		return i1/i2
	})

	divide(1,0)

	println("main end...")
}

func Divide(fn func(int, int) int) func(int, int) int {
	return func(x, y int) int {
		defer func() {
			if e := recover(); e != nil {
				fmt.Printf("this a %v\n", e)
			}
		}()

		return fn(x, y)
	}
}
