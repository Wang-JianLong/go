// 为 int 定义一个别名类型 Day，
// 定义一个字符串数组它包含一周七天的名字，为类型 Day 定义 String() 方法，
// 它输出星期几的名字。使用 iota 定义一个枚举常量用于表示一周的中每天（MO、TU...）。
package main

import "fmt"

const (
	Monday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

type Day int

func (d Day) String() string {
	SS := [7]string{"Monday", "Thuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	return SS[d]
}

func main() {
	day := Day(Monday)
	fmt.Printf("day: %v\n", day) // day: Monday
}
