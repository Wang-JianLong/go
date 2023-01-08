package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	bf := bufio.NewReader(os.Stdin)
	s, err := bf.ReadString('S')
	if err == nil{
		size:=0
		lineSize := 0
		for i := 0; i < len(s); i++ {
			if  s[i] != '\r'|| s[i] != '\n'{
				size++
			}
			if s[i] == '\n' {
				lineSize ++
			}
			
		
		}

		// 单词
		fieldSize := len(strings.Fields(s))
		
		fmt.Printf("size: %v\n", size)
		fmt.Printf("fieldSize: %v\n", fieldSize)
		fmt.Printf("lineSize: %v\n", lineSize)
	}
}