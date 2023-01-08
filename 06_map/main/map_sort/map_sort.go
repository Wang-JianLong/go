package main

import (
	"fmt"
	"sort"
)

var (
	barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}
)

func main() {
	// map 默认是无序的，不管是按照 key 还是按照 value 默认都不排序
	// 想为 map 排序，
	// 需要将 key（或者 value）`拷贝`到一个切片，再对切片排序
	keys := make([]string,0,len(barVal))
	vals := make([]int,0,len(barVal))
	for k,v := range barVal{
		keys = append(keys, k)
		vals = append(vals, v)
	}

	sort.Strings(keys)
	sort.Ints(vals)

	fmt.Printf("keys: %v\n", keys)
	fmt.Printf("vals: %v\n", vals)

	/*
	keys: [alpha bravo charlie delta echo foxtrot golf hotel indio juliet kili lima]
	vals: [12 16 23 34 34 43 56 56 65 87 87 98]
	*/
}

func sorted(){
	fmt.Println("unsorted:")
    for k, v := range barVal {
        fmt.Printf("Key: %v, Value: %v / ", k, v)
    }
    keys := make([]string, len(barVal))
    i := 0
    for k, _ := range barVal {
        keys[i] = k
        i++
    }
    sort.Strings(keys)
    fmt.Println()
    fmt.Println("sorted:")
    for _, k := range keys {
        fmt.Printf("Key: %v, Value: %v / ", k, barVal[k])
    }

}