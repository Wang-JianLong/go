package main

import (
	"fmt"
	"io/ioutil"
)

/*
请给这个结构编写一个 save 方法，
	将 Title 作为文件名、Body 作为文件内容，写入到文本文件中。

再编写一个 load 函数，
	接收的参数是字符串 title，
	该函数读取出与 title 对应的文本文件。请使用 *Page 做为参数，
	因为这个结构可能相当巨大，我们不想在内存中拷贝它。请使用 ioutil 包里的函数（
*/
type Page struct {
    Title string
    Body  []byte
}

func (p Page) String()string{
	return "Title;Body;\n"+p.Title+";"+string(p.Body)+";\n"
}

func (p Page) save() (e error){
	e = ioutil.WriteFile(p.Title,p.Body,0666)
	return
}

func (p *Page )Load(title string)(e error){
	p.Title = title
	p.Body,e = ioutil.ReadFile(p.Title)
	return 
}

func main(){
	page := &Page{"A",[]byte{'1','2','3'}}
	page.Title = "wjl.bytes"
	page.save()
	

	p1 := &Page{Title: page.Title}
	p1.Load(p1.Title)
	fmt.Printf("p1: %v\n", p1)
}