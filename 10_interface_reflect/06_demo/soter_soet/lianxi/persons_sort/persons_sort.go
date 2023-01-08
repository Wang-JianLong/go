package main

import (
	"fmt"
	"strings"
)

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func Sort(data Sorter) {
	for pass := 1; pass < data.Len(); pass++ {
		for i := 0; i < data.Len()-pass; i++ {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
			}
		}
	}
}

/*
定义一个结构体 Person，

	它有两个字段：firstName 和 lastName，为 []Person 定义类型 Persons 。

让 Persons 实现 Sorter 接口并进行测试。
*/
type Person struct {
	firstName, lastName string
}

type PersonArray []Person

func (p PersonArray) Len() int {
	return len(p)
}

func (p PersonArray) Less(i, j int) bool {
	p1,p2 := p[i],p[j]
	less := strings.Compare(p1.firstName,p2.firstName)
	if less == 0{
		return strings.Compare(p1.lastName,p2.lastName) <= 0
	}
	return  less < 0
}

func (p PersonArray) Swap(i, j int){
	p[i],p[j] = p[j],p[i]
}

func main() {
	ps := PersonArray{Person{"asda", "ada"}, Person{"dasda", "adasdsa"}, Person{"dasdas", "dadasgdg"}, Person{"asda", "fdghfdghd"}}

	Sort(ps)

	fmt.Printf("ps: %v\n", ps)
}