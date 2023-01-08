package main

import (

	"encoding/json"
	"fmt"
	"os"
)



func main() {
	addr := &Address{"one", "china", "xian"}
	v := &VCard{"petter", "paker", []*Address{addr}, "yes"}

	bys,err := json.Marshal(v)
	if err != nil{
		//...
	}
	fmt.Printf("string(bys): %v\n", string(bys))

	file,err := os.OpenFile("jsonText.json",os.O_WRONLY|os.O_CREATE,0666)

	encode := json.NewEncoder(file)
	encode.Encode(v)
	file.Close()

	v1 := &VCard{}
	json.Unmarshal(bys,v1)
	fmt.Printf("v1: %v\n", v1)
	fmt.Printf("v1.Addresses: %v\n", v1.Addresses[0])
	// 不匹配会被忽略掉



}