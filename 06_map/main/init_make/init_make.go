package main

import "fmt"

/*
map 是一种特殊的数据结构：一种元素对（pair）的无序集合，
	pair 的一个元素是 key，对应的另一个元素是 value，所以这个结构也称为关联数组或字典。
	这是一种快速寻找值的理想结构：给定 key，对应的 value 可以迅速定位。
map 这种数据结构在其他编程语言中也称为字典（Python）、hash 和 HashTable 等。
*/
func main() {
	// map是引用类型
	// 1. 声明 时不需要知道map的长度 是可以动态增长的 未初始化的值为nil
	// var m map[string]int;
	/*
		Key
		key 可以是任意可以用 == 或者！= 操作符比较的类型，比如 string、int、float。
		所以切片和结构体不能作为 key
		(译者注：含有数组切片的结构体不能作为 key，只包含内建类型的 struct 是可以作为 key 的），
			但是指针和接口类型可以。
		如果要用结构体作为 key 可以提供 Key() 和 Hash() 方法，
		这样可以通过结构体的域计算出唯一的数字或者字符串的 key。

		Value

	*/

	var mapLit map[string]int
	var mapAssigned map[string]int

	// 直接初始化
	mapLit = map[string]int{"one": 1, "two": 2}
	// make初始化
	mapCreated := make(map[string]float64)

	mapAssigned = mapLit

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14

	mapAssigned["tow"] = 3

	fmt.Printf("mapLit[\"one\"]: %v\n", mapLit["one"])
	fmt.Printf("mapCreated[\"key1\"]: %v\n", mapCreated["key1"])
	fmt.Printf("mapAssigned[\"tow\"]: %v\n", mapAssigned["tow"])
	fmt.Printf("mapLit[\"tow\"]: %v\n", mapLit["tow"])
	fmt.Printf("mapLit[\"ten\"]: %v\n", mapLit["ten"])
	/*
	mapLit["one"]: 1
	mapCreated["key1"]: 4.5
	mapAssigned["tow"]: 3
	mapLit["tow"]: 3
	mapLit["ten"]: 0
	*/

	/*
	不要使用 new，永远用 make 来构造 map
	注意 如果你错误的使用 new () 分配了一个引用对象，你会获得一个空引用的指针，
	相当于声明了一个未初始化的变量并且取了它的地址：
	*/

	// m := new(map[string]int)
	// m["k1"] = 1
	// fmt.Printf("m: %v\n", m)
	// invalid operation: cannot index m (variable of type *map[string]int)
}
