package main

import "fmt"

func main() {
	m := map[string]string{"1": "One", "2": "Two"}
	// 注意 map 不是按照 key 的顺序排列的，也不是按照 value 的序排列的
	for k, v := range m {
		fmt.Printf("k:%v,v:%v\n",k,v)
	}
}

// 创建一个 map 来保存每周 7 天的名字，将它们打印出来并且测试是否存在 Tuesday 和 Hollyday
func Week(){
	week := map[int]string{
		1 :"Monday",
		2 :"Tuesday",
		3 :"Wednesday",
		4 :"Thursday",
		5 :"Friday",
		6 :"Saturday",
		7 :"Sunday",
	}

	for k,v := range week{
		fmt.Printf("k:%v,v:%v\n",k,v)
	}

	
}
