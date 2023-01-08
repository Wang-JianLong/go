package main

import "fmt"

func main() {
	/*
		和数组不同，map 可以根据新增的 key-value 对动态的伸缩，因此它不存在固定长度或者最大限制。
		但是你也可以选择标明 map 的初始容量 capacity

		当 map 增长到容量上限的时候，如果再增加新的 key-value 对，map 的大小会自动加 1。
		所以出于性能的考虑，对于大的 map 或者会快速扩张的 map，
		即使只是大概知道容量，也最好先标明。

	*/
	// m := make(map[int]string,100)
	noteFrequency := map[string]float32{
		"C0": 16.35, "D0": 18.35, "E0": 20.60, "F0": 21.83,
		"G0": 24.50, "A0": 27.50, "B0": 30.87, "A4": 440,
	}

	fmt.Printf("noteFrequency: %v\n", noteFrequency)

	m := make(map[int]string,1)
	m[1] = "1"
	m[2] = "2"
	fmt.Printf("m: %v\n", m)

}