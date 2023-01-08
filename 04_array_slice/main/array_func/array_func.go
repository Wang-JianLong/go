package main

/*
把一个大数组传递给函数会消耗很多内存。有两种方法可以避免这种现象：
    传递数组的指针
    使用数组的切片
*/
func main(){

}

func f1(a *[10]float64) (f float64 ){
	for i := 0; i < len(*a); i++ {
			f += (*a)[i];
	}
	// 但这在 Go 中并不常用，通常使用切片
	return
}