package main

import "fmt"

func main() {
	s := []string{"1", "2", "3","4"}
	res, ok := RemoveSlice(s, 1, 2)
	if ok {
		fmt.Printf("res: %v\n", res)
	}
}

// 编写一个函数 RemoveStringSlice 将从 start 到 end 索引的元素从切片 中移除

func RemoveSlice(s []string, start, end int) (res []string, ok bool) {
	if start > end || start < 0{
		ok, res = false, s
		return
	}
	if start == end {
		ok, res = true, s
		return
	}

	ok, res = true, append(s[:start],s[end + 1 :]...)
	return
}