package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	inputFile, fileOpenErr := os.Open("D:/Code/Go/src/11_bufio/02_read_file/wjl.txt")
	if fileOpenErr != nil {
		fmt.Print("error %v", fileOpenErr)
		return
	}
	// 开启关闭资源
	defer inputFile.Close()

	bys, err := io.ReadAll(inputFile)
	if err == nil {
		fmt.Printf("string(bys): %v\n", string(bys))
	}
	fmt.Println("******************************")
	bs,_:= ioutil.ReadFile("D:/Code/Go/src/11_bufio/02_read_file/wjl.txt")
	fmt.Printf("string(bs): %v\n", string(bs))
	fmt.Println("******************************")
	// 读取资源 // 文件流被读取过就空了
	// newFile,_ := os.Open(inputFile.Name())
	bf := bufio.NewReader(inputFile)
	for {
		str, readErr := bf.ReadString('\n')
		fmt.Printf("str: %v\n", str)
		if readErr == io.EOF {
			return
		}
	}

}
