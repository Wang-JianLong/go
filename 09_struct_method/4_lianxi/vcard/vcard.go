package main

import (
	"fmt"
	"time"
)

/*
定义结构体 Address 和 VCard，
后者包含一个人的名字、地址编号、出生日期和图像，
试着选择正确的数据类型。
构建一个自己的 vcard 并打印它的内容。
*/
func main(){
	var t time.Time = time.Now()
	tp := &t
	vcard := VCard{"zhangsan",&Address{"China"},tp,"http://images...."}
	fmt.Printf("vcard: %v\n", vcard)
	// vcard: {zhangsan {China} 2023-01-01 17:39:29.0286332 +0800 CST m=+0.002071101 http://images....}
}

type Address struct{
	name string
}

type VCard struct {
	name string
	address * Address
	Birthday * time.Time
	image string
}