package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func Sort(data Sorter) {
	for pass := 1; pass < data.Len(); pass++ {
		for i := 0; i < data.Len()-pass; i++ {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
			}
		}
	}
}

func IsSorted(data Sorter) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}

type Float64Array []float64

func (f Float64Array) Len() int {
	return len(f)
}

func (f Float64Array) Less(i, j int) bool {
	return f[i] < f[j]
}

func (f Float64Array) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f Float64Array) Fill() [10]float64{
	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)
	var rands [10]float64;
	for i := range rands{
		rands[i] = 100 * rand.Float64()
	}
	return rands
}
		



func main() {
	data := []float64{1.2, 1.1, 1.12, 1.32, 34, 0.001, 0.1, 0.3}
	// 强转
	array := Float64Array(data)


	Sort(array)

	fmt.Printf("data: %v\n", data)

	fmt.Printf("array.Fill(): %v\n", array.Fill())
}
