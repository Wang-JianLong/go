package main

import (
	"flag"
	"os"
)

// 定义具有指定名称、默认值和用法字符串的 bool 标志。返回值是存储标志值的布尔变量的地址。
var NewLine = flag.Bool("n", false, "print newline") // echo -n flag, of type *bool

// 空格和换行
const (
	Space   = " "
	Newline = "\n"
)

func main() {
	flag.PrintDefaults()
	flag.Parse() //解析从操作系统解析命令行标志。参数[1：]。必须在定义所有标志之后和程序访问标志之前调用。

	var s string = ""
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += "   "
			if *NewLine{
				s += Newline
			}
		}
		s += flag.Arg(i)
	}
	
	os.Stdin.WriteString(s)
}