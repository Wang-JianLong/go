package main

import "math"

type Point struct {
	x, y float64
}

func (point *Point) Abs() (res float64) {
	x := point.x
	y := point.y
	res = math.Sqrt(x*x + y*y)
	return
}

type Polar struct {
	point   *Point
	z, Sqrt float64
}

func (p *Polar) Scale(f float64) {
	p.point.x *= f
	p.point.y *= f
	p.z *= f
}

func main() {

}
