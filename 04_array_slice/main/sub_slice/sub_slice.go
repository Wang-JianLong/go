package main

import "fmt"

func main() {
	// 申请了一份内存空间 底层数组 一共十个 针对此切片是5个
	s := make([]int, 5, 10)
	fmt.Printf("s: %v\n", s)
	fmt.Printf("len(s): %v\n", len(s))
	fmt.Printf("cap(s): %v\n", cap(s))

	for i := 0; i < 5; i++ {
		s[i] = 1;
	}

	//s[5] = 1 // 此时若是添加第六个元素 那么 [死机：运行时错误：索引超出范围[5]，长度为5]

	// 创建一个子切片 设定从母切片的2[第三个]开始 4[第五个]结束
	sub := s[2:4]
	fmt.Printf("sub: %v\n", sub)
	// 那么len是预料之中的
	fmt.Printf("len(sub): %v\n", len(sub)) // 2 
	// 那么容量8？
	fmt.Printf("cap(sub): %v\n", cap(sub)) // 8
  
	// 再来一个子元素
	sub2 := s[3:5]
	fmt.Printf("sub2: %v\n", sub2)
	fmt.Printf("len(sub2): %v\n", len(sub2)) // 2
	fmt.Printf("cap(sub2): %v\n", cap(sub2)) // 7

	// 结论: 子切片的cap容量 = 所依赖的母切片的cap - 当前子切片的首索引位置

	sub[0] = 111;
	sub2[0] = 1111;
	fmt.Printf("s: %v\n", s) //s: [1 1 111 1111 1]

	sub = append(sub,1)
}
