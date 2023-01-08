package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"os"
)

func main() {
	fileName := "D:/Code/Go/src/11_bufio/02_read_file/read_zip/aaa.txt.gz"
	reader, err := os.Open(fileName)
	if err != nil {
		panic("文件异常")
	}
	defer reader.Close()
	var bf *bufio.Reader

	fz, fzerr := gzip.NewReader(reader)
	if fzerr != nil {
		bf = bufio.NewReader(reader)
	} else {
		bf = bufio.NewReader(fz)
	}

	for {
		line, err := bf.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Printf("line: %v\n", line)
	}

}
