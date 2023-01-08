package main

import "fmt"

func main() {
	m := map[int]func() int{
		1: func() int { return 1 },
		2: func() int { return 2 },
	}

	fmt.Printf("m[1](): %v\n", m[1]()) // 1
	
	m2 := make(map[int]func(k int,v string) string)

	m2[1] = func(k int, v string) string {
		return v
	} 

	m2[2] = func(k int, v string) string {
		return v + "2"
	}

	fmt.Printf("m2: %v\n", m2) //m2: map[1:0x2fff60 2:0x2fff80]
}