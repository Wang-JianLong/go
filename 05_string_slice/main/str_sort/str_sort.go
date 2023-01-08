package main

import (
	"fmt"
	"sort"
)

// 标准库提供了 sort 包来实现常见的搜索和排序操作
func main(){
	// 排序int切片
	i := []int{23,435,78,1,34,56,12}
	sort.Ints(i)
	fmt.Printf("i: %v\n", i)
	// 是否排序
	// ntsAreSort 报告切片 x 是否按递增顺序排序。
	isSort := sort.IntsAreSorted(i)
	fmt.Printf("isSort: %v\n", isSort)

	i2 := []int{1,2,3}
	fmt.Printf("sort.IntsAreSorted(i2): %v\n", sort.IntsAreSorted(i2))

	s := []string{"A","C","B"}
	sort.Strings(s)
	fmt.Printf("s: %v\n", s)

	// SearchString 在排序的字符串切片中搜索 x，并返回 Search 指定的索引。返回值是 x 不存在时插入 x 的索引（可以是 len（a））。
	//切片必须按升序排序。
	fmt.Printf("sort.SearchStrings(s, \"C\"): %v\n", sort.SearchStrings(s, "C"))
}