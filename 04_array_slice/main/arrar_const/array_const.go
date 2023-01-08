package main

import "fmt"

// 几何点（或者数学向量）是一个使用数组的经典例子。为了简化代码通常使用一个别名：
type Vector3D [3]int

var vec Vector3D

func main() {
	// 如果数组值已经提前知道了，那么可以通过 数组常量 的方法来初始化数组，而不用依次使用 []= 方法（所有的组成元素都有相同的常量语法）
	arr1 := [5]string{1: "Jack", 4: "Deg", 2: "Jerry"}
	fmt.Printf("arr1: %v\n", arr1) // arr1: [ Jack Jerry  Deg]

	arr2 := [10]int{1, 2, 3}
	fmt.Printf("arr2: %v\n", arr2) // arr2: [1 2 3 0 0 0 0 0 0 0]

	arr3 := [...]int{1, 2, 3}
	fmt.Printf("arr3: %v\n", arr3) // arr3: [1 2 3]

}
