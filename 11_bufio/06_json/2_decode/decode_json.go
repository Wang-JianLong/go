package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// json 包使用 map[string]interface{} 和 []interface{} 储存任意的 JSON 对象和数组；其可以被反序列化为任何的 JSON blob 存储到接口值中。
func main() {
	b := []byte(`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia"]}`)
	var f interface{}
	json.Unmarshal(b,&f)
	
	if m,ok := f.(map[string]interface{});ok{
		fmt.Printf("m: %v\n", m)
	}
	var fa FamilyMember
	json.Unmarshal(b,&fa) // 实际上分出了一个新的切片
	fmt.Printf("fa: %v\n", fa)

	// 解码文件
	file,err := os.Open("jsonText.json")
	if err != nil{
		// ...
	}
	// buf...
	decode := 	json.NewDecoder(file)
	var vcard VCard
	decode.Decode(&vcard)
	fmt.Printf("vcard: %v\n", vcard)

}
// 如果事先知道json的结构类型
type FamilyMember struct {
    Name    string
    Age     int
    Parents []string
}
type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}