package main
import "fmt"
func main(){
	s := []int{1,2,3,4,5,6,7,8,9,10}
	s = s[:]
	s = Filter(s,func (num int) bool {
		return num % 2 == 0
	})
	fmt.Printf("%v\n",s)
	fmt.Printf("s: %#v\n", s)
}
/*
使用高阶函数对一个集合进行过滤：s 是前 10 个整数的一个切片。
建立一个函数 Filter，它接受 s 作为第一参数，
fn func (int) bool 作为第二参数，并返回满足函数 fn 的 s 元素的切片（使其为真）。
用 fn 测试整数是否为偶数
*/
// TODO
func Filter(s []int,fun func(int) bool) []int {
	res := make([]int,1,1)
	isOne := true
	for i,v := range s {
		// 拷贝第一个参数
		if fun(v) && isOne {
			copy(res,s[i:i+1])	
			isOne = false
		}else if fun(v){
			res = append(res, s[i])
		}
	}
	return res
}