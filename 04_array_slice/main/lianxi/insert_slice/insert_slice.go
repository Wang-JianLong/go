package main

import "fmt"

func main() {
	from := []int{1, 2, 3}
	to := make([]int, 4)
	res, ok := InsertStringSlice(to, from, 3)
	if ok {
		fmt.Printf("res: %v\n", res)
	}
}

// 编写一个函数 InsertStringSlice 将切片插入到另一个切片的指定位置。
// 既然是插入 那么就是指定下标前面
func InsertStringSlice[T []int](to T, from T, index int) (res []int, ok bool) {
	toLen := len(to)

	if index < toLen { // 判断 指定下标位置是否合法
		before_s, after_s := to[:index], to[index:]
		res = append(before_s, from...) // 展开表达式
		res = append(res, after_s...)
		ok = true
	} else {
		res = make([]int, len(from)+index+toLen)
	}

	return
}
