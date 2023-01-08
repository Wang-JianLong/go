package main

import "fmt"

/*
如果 map 中不存在 key1，val1 就是一个值类型的空值(初始值)

现在我们没法区分到底是 key1 不存在还是它对应的 value 就是空值。
*/
func main(){
	m := map[int]string{
		1:"zs",
		2:"lisi",
	}

	v,isPrent := m[1]
	fmt.Printf("v: %v\n", v) // “zs”
	fmt.Printf("isPrent: %v\n", isPrent) // true

	v,isPrent = m[3]
	fmt.Printf("v: %v\n", v) // 
	fmt.Printf("isPrent: %v\n", isPrent) // false

	// 删除k
	delete(m,1) // 就算k不存在，也不会产生错误
	delete(m,3)

	fmt.Printf("m: %v\n", m)
}