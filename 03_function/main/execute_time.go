package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for1000();
	end := time.Now()

	executeTime := end.Sub(start)

	fmt.Printf("end: %v\n", executeTime)
}

func for1000(){
	for i:=10000;i>0;i--{

	}
}