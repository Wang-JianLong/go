package main

import "fmt"

/*
个接口类型的变量 varI 中可以包含任何类型的值，
	必须有一种方式来检测它的 动态 类型，即运行时在变量中存储的值的实际类型。
	在执行过程中动态类型可能会有所不同，但是它总是可以分配给接口变量本身的类型。
	通常我们可以使用 类型断言 来测试在某个时刻 varI 是否包含类型 T 的值

	varI 必须是一个接口变量，否则编译器会报错：invalid type assertion: varI.(T) (non-interface type (type of varI) on left) 。

	类型断言可能是无效的，虽然编译器会尽力检查转换是否有效，但是它不可能预见所有的可能性。
	如果转换在程序运行时失败会导致错误发生。更安全的方式是使用以下形式来进行类型断言：

*/

type I interface{
	I()
}

type Ii struct{

}

func (i *Ii) I(){

}

func main(){
	var i I = &Ii{}

	if v,ok:=i.(*Ii) ;ok{
		fmt.Printf("v: %v\n", v) // v: &{}
		fmt.Printf("ok: %v\n", ok) // ok: true
		// 执行逻辑
		// 应该总是使用上面的方式来进行类型断言。
	}
	
}