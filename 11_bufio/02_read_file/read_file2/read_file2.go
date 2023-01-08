package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("D:\\Code\\Go\\src\\11_bufio\\02_read_file\\read_file2\\abcd.txt")
	if err != nil{
		panic("文件为空")
	}
	var co1, co2, co3 []string
	for {
		var v1, v2, v3 string
		_, err := fmt.Fscanln(io.Reader(file), &v1, &v2, &v3)
		if err != nil {
			break
		}

		co1 = append(co1,v1)
		co2 = append(co2,v2)
		co3 = append(co3,v3)

	}
	defer file.Close()

	fmt.Printf("co1: %v\n", co1)
	fmt.Printf("co2: %v\n", co2)
	fmt.Printf("co3: %v\n", co3)
}
