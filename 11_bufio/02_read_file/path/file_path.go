package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	str := filepath.Base("D:\\Code\\Go\\src\\11_bufio\\02_read_file\\read_file2\\abcd.txt")
	fmt.Printf("str: %v\n", str)
}