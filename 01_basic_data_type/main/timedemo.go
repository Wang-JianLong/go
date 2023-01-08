package main

import (
	"fmt"
	"time"
)

func main() {
	var timer time.Time = time.Now()
	fmt.Println(timer)
	fmt.Println(timer.Day())
}
