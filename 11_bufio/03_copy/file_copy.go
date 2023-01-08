package main

import (
	"io"
	"os"
)

func main() {
	CopyFile("target.txt","source.txt")
	
	target := "D:/Downloads/NormalFile/video/hmn-174-C.mp4"
	CopyFile(target,"h.mp4")
}

func CopyFile(target, source string) (int64, error) {
	// 获取输入流
	src, err := os.Open(target)
	if err != nil{
		return 0,err
	}
	defer src.Close()
	// 获取输出流
	dst,err  := os.OpenFile(source,os.O_WRONLY|os.O_CREATE,0666)
	if err != nil{
		return 0,err
	}
	defer dst.Close()

	// 注意 defer 的使用：当打开目标文件时发生了错误，
	// 那么 defer 仍然能够确保 src.Close() 执行。
	// 如果不这么做，文件会一直保持打开状态并占用资源。
	// defer 后进先出原则
	return io.Copy(dst,src)
}