package main

import "fmt"

// import "fmt"
// 测试某一个结构体是否实现了接口

type Inr interface {
    In()
}

type InIm struct {
    
}

func (i *InIm) In(){
    
}

/*

type Stringer interface {
    String() string
}

if sv, ok := v.(Stringer); ok {
    fmt.Printf("v implements String(): %s\n", sv.String()) // note: sv, not v
}
*/


func main(){
    fmt.Printf("\"a\": %v\n", "a")
}
