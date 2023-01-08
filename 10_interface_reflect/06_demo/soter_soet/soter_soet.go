package main

import (
	"fmt"
	"sort"
)

/*
一个很好的例子是来自标准库的 sort 包，要对一组数字或字符串排序，只需要实现三个方法

// Sort 函数接收一个接口类型参数：Sorter ，它声明了这些方法：
type Sorter interface {
    Len() int // 长度
    Less(i, j int) bool 比较
    Swap(i, j int) // 交换
}

现在如果我们想对一个 int 数组进行排序，
所有必须做的事情就是：为数组定一个类型并在它上面实现 Sorter 接口的方法

*/
type IntArry []int

func (i IntArry) Len() int {
	return len(i)
}

func (a IntArry) Less(i, j int) bool {
	return a[i] < a[j]
}

func (a IntArry) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	// 包装
	array := IntArry(data)
	sort.Sort(array)
	
	// data: [-5467984 -784 0 0 42 59 74 238 905 959 7586 7586 9845]
	fmt.Printf("data: %v\n", data)
	// 注意 data 是 slice 本身就是引用 所以IntArray不需要 *  本来就是指针

}
