package main

/*
假设你有一个 myType 类型的数据切片，你想将切片中的数据复制到一个空接口切片中，类似：

var dataSlice []myType = FuncReturnSlice()
var interfaceSlice []interface{} = dataSlice

可惜不能这么做，编译时会出错：cannot use dataSlice (type []myType) as type []interface { } in assignment。

原因是它们俩在内存中的布局是不一样的（参考 官方说明 https://golang.org/doc/go_spec.html）

必须使用 for-range 语句来一个一个显式地复制：

var dataSlice []myType = FuncReturnSlice()
var interfaceSlice []interface{} = make([]interface{}, len(dataSlice))
for i, d := range dataSlice {
    interfaceSlice[i] = d
}
*/
func main() {
	ss := make([]string, 0)
	ss = append(ss, "a")
	ss = append(ss, "b")
	ss = append(ss, "b")
	// inters declared but not used
	// var inters  []interface{} = ss
}
