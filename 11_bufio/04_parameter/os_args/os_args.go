package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 1{
		fmt.Print(strings.Join(os.Args[1:]," 你好"))
		 // go run .\os_args.go os_args John Bill Marc Luke
		 // 运行 os_args 你好John 你好Bill 你好Marc 你好Luke

	}
}