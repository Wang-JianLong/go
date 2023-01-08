package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 从控制台读取
	var inputReadr *bufio.Reader = bufio.NewReader(os.Stdin)
	inpu, err := inputReadr.ReadString('%')
	if err == nil{
		fmt.Printf("inpu: %v\n", inpu)		
	}

		
}
