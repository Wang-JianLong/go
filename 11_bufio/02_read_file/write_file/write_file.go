package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileName := "D:/Code/Go/src/11_bufio/02_read_file/write_file/123.txt"
	// OpenFile 是广义的开放调用;大多数用户将改用“打开”或“创建”。它打开带有指定标志 （O_RDONLY 等）的命名文件。如果该文件不存在，并且传递了 O_CREATE 标志，则使用模式 perm 创建该文件（在 umask 之前）。
	outputFile, outputError := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)

	// 在读文件的时候，文件的权限是被忽略的，所以在使用 OpenFile 时传入的第三个参数可以用 0。而在写文件时，不管是 Unix 还是 Windows，都需要使用 0666。

	if outputError != nil{
		panic("Error 文件异常")
	}

	defer outputFile.Close()
	// 放入包装
	bf := bufio.NewWriter(outputFile)
	bf.WriteString("我是这个小区最牛逼的人")
	
	bf.Flush()

	fmt.Fprintf(outputFile,"Hello\n123")

}