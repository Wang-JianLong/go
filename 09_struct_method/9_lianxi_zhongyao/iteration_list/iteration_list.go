// 下面这段代码有什么错？
package main


// import "container/list"

// 无法在非本地类型列表上定义新方法。
// func (p *list.List) Iter() {
//     // ...
// }

// func main() {
//     lst := new(list.List)
//     for _= range lst.Iter() {
//     }
// }

// 类型和作用在它上面定义的方法必须在同一个包里定义，
// 这就是为什么不能在 int、float 或类似这些的类型上定义方法。
// 试图在 int 类型上定义方法会得到一个编译错误 cannot define new methods on non-local type int

// 类型在其他的，或是非本地的包里定义，在它上面定义方法都会得到和上面同样的错误。

// 但是有一个间接的方式：
// 可以先定义该类型（比如：int 或 float）的别名类型，
// 然后再为别名类型定义方法。或者像下面这样将它作为匿名类型嵌入在一个新的结构体中。
// 当然方法只在这个别名类型上有效。


import (
    "fmt"
    "time"
)

// 继承 + 扩展 + 包装 ？
type myTime struct {
    time.Time //anonymous field
}

func (t myTime) first3Chars() string {
    return t.Time.String()[0:3]
}

func main() {
    m := myTime{time.Now()}
    // 调用匿名Time上的String方法
    fmt.Println("Full time now:", m.String())
    // 调用myTime.first3Chars
    fmt.Println("First 3 chars:", m.first3Chars())
}

/* Output:
Full time now: Mon Oct 24 15:34:54 Romance Daylight Time 2011
First 3 chars: Mon
*/