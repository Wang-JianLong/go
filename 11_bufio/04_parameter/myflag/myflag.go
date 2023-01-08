package main

import (
	"flag"
	"fmt"
)

var name, password string

func main() {
	flag.StringVar(&name, "u", "root", "用户名")
	flag.StringVar(&password, "p", "", "密码")
	flag.Parse()
	// .\myflag.exe -u wjl -p 123456
	if password == "123456" {
		fmt.Printf("user:" + name + "; p:" + password + "  登录成功")
	} else {
		fmt.Println("FAIDE!!!")
	}
}
