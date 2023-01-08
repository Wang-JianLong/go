package main

import (
	"fmt"
)

type Base struct{}

func (Base) Magic() {
	fmt.Println("base magic")
}

func (self Base) MoreMagic() {
	self.Magic()
	self.Magic()
}

type Voodoo struct {
	Base
}

func (Voodoo) Magic() {
	fmt.Println("voodoo magic")
}

// 不复写 就是调用Base(父类)的两个Magic  self之前指向的是Base实例
func (self Voodoo) MoreMagic() {
	self.Magic()
	self.Magic()
}

func main() {
	v := new(Voodoo)
	v.Magic()
	v.MoreMagic()
}
