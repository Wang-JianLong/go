package main

func main() {
	println(minSlice([]int{3, 2, -7, 1, 0, -3}))
	println(maxSlice([]int{1, 2, 3, 4, 7, 0, 4, 56, 45}))
}

/*
写一个 minSlice 方法，传入一个 int 的切片并且返回最小值，再写一个 maxSlice 方法返回最大值
*/

func minSlice(is []int) (min int) {
	min = is[0]
	for _, v := range is {
		if v < min {
			min = v
		}
	}
	return
}

func maxSlice(is []int) (max int) {
	max = is[0]
	for _, v := range is {
		if v > max {
			max = v
		}
	}
	return
}
