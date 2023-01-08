// package main
package point
import (
	"fmt"
	"math"
)

/*
	使用坐标 X、Y 定义一个二维 Point 结构体。
	同样地，对一个三维点使用它的极坐标定义一个 Polar 结构体。
	实现一个 Abs() 方法来计算一个 Point 表示的向量的长度，
	实现一个 Scale 方法，它将点的坐标乘以一个尺度因子
	（提示：使用 math 包里的 Sqrt 函数）。
*/
type Point struct{
	x,y float64
}

func Abs(point *Point)(res float64){
	x := point.x
	y := point.y
	res  = math.Sqrt(x*x + y*y)
	return 
}

type Polar struct {
	point *Point
	z float64
}
func Scale(p *Polar,f float64){
	p.point.x *= f
	p.point.y *= f
	p.z *= f
}

// func main(){
// 	point := &Point{123,453}
// 	res := Abs(point)
// 	fmt.Printf("res: %v\n", res)

// 	polar := &Polar{point,100}
// 	Scale(polar,100)
// 	fmt.Printf("point: %v\n", point)

// 	//	res: 469.40174690769953
// 	//	point: &{12300 45300}
	

// }




func abc(){
	fmt.Printf("\"abc\": %v\n", "abc")
}