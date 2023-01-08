package main

import "fmt"

func main() {
	f := fib()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", f())
	}
}

func fib() func() int {
	first := 0
	next := 1
	return func() int {
		defer func() {
			first, next = next, (first + next)
		}()
		return first
	}
}
