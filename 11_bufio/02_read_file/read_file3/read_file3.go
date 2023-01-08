package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	title string
	 price float64
	  quantity int
}

func main() {

	file, err := os.Open("D:/Code/Go/src/11_bufio/02_read_file/read_file3/aaa.txt")
	if err != nil {
		panic("文件打开异常")
	}
	defer file.Close()
	var ps []Product = make([]Product, 0)

	fileReader := bufio.NewReader(file)

	for {
		str, err := fileReader.ReadString('\n')
		str = str [: len(str) - 2] // win 换行是 /r/n 会导致最后一个转换数字失败
		if err == io.EOF {
			fmt.Println("没有文件可读取了...")
			break
		}
		strs := strings.Split(str, ";")
		it,_ := strconv.Atoi(strs[2])
		fl,_:= strconv.ParseFloat(strs[1],64)
		ps = append(ps, Product{strs[0],fl,it})
		// fmt.Printf("str: %v\n", strs)
	}
	fmt.Printf("ps: %v\n", ps)

}
