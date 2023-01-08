package main

func main() {

}

/*
编写一个函数 Sum，其参数是一个包含 4 个浮点数的数组 arrF，并返回数组中所有数字的总和。如果使用切片而不是数组，代码要如何修改？该函数的切片形式适用于不同长度的数组！
*/
func Sum(arrF [4]float64) (result float64) {
	for _, v := range arrF {
		result += v
	}
	return
}

/*
	编写一个函数 SumAndAverage，将这两个未命名的 int 和 float32 类型的变量返回
*/
func SumAndAverage(a, b int, c float64) (sum, average float64) {
	// 返回两个 int 和 float32 类型的未命名变量的和与平均值
	sum, average = float64(a)+float64(b)+c, sum/3
	return
}
