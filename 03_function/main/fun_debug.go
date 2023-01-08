package main

import (
	"fmt"
	"log"
	"runtime"
)

func main() {
	// array := [...]int {1,2,3};
	// fmt.Printf("array: %v\n", array)

	// s := []int{1,2,3,4,5}
	// fmt.Printf("s: %T\n", s)
	test2()
}

func test1(){
	where := func() {
		pc, file, line, ok := runtime.Caller(1)
		fmt.Printf("pc: %v file:%v line:%v ok:%v\n", pc, file, line, ok)
	}

	where()
	where()
	where()
	/*
		pc: 8453473 file:d:/Code/Go/src/03_function/main/fun_debug.go line:19 ok:true
		pc: 8453489 file:d:/Code/Go/src/03_function/main/fun_debug.go line:20 ok:true
		pc: 8453505 file:d:/Code/Go/src/03_function/main/fun_debug.go line:21 ok:true
	*/
}

func test2(){
	log.SetFlags(log.Llongfile)
	log.Print("this is a debug");// d:/Code/Go/src/03_function/main/fun_debug.go:36: this is a debug
}
