package main

import (
	"fmt"
	"math"
)

type Shaper interface {
    Area() float32
}

type shaper interface {
	say(string2 string)
}
type shape struct {
}

func (a *shape) say(string2 string) {
	fmt.Print("i am ", string2)
}

type Circle struct{
    radius float32
    shape
}

func (c *Circle) Area()float32{
    return math.Pi * c.radius
}
func main() {

    // r := Rectangle{5, 3} // Area() of Rectangle needs a value
    // q := &Square{5}      // Area() of Square needs a pointer
    // // shapes := []Shaper{Shaper(r), Shaper(q)}
    // // or shorter
    // shapes := []Shaper{r, q}
    // fmt.Println("Looping through shapes for area ...")
    // for n, _ := range shapes {
    //     fmt.Println("Shape details: ", shapes[n])
    //     fmt.Println("Area of this shape is: ", shapes[n].Area())
    // }

    c:= &Circle{5,shape{}}
    fmt.Printf("c.Area(): %v\n", c.Area())
    c.say("zhangsna")
}