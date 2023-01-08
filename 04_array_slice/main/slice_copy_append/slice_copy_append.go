package main

import "fmt"

/*
如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来
*/
// func main(){
// 	from := []int{1,2,3}
// 	to := make([]int,10,15)
// 	n := copy(to,from)
// 	fmt.Printf("from: %v\n", from)
// 	fmt.Printf("to: %v\n", to)
// 	fmt.Printf("n: %v\n", n)
// }

/*
func append(s[]T, x ...T) []T 其中 append 方法将 0 个或多个具有相同类型 s 的元素追加到切片后面并且返回新的切片；追加的元素必须和原切片的元素同类型。如果 s 的容量不足以存储新增元素，append 会分配新的切片来保证已有切片元素和新增元素的存储。因此，返回的切片可能已经指向一个不同的相关数组了。append 方法总是返回成功，除非系统内存耗尽了。

如果你想将切片 y 追加到切片 x 后面，只要将第二个参数扩展成一个列表即可：x = append(x, y...)
*/

func main() {
	s1 := []int{1, 2, 3}
	s1 = append(s1, 1, 2, 3, 4)
	fmt.Printf("s1: %v\n", s1)
	
}

func Append(sl []byte,data...[]byte){

}
