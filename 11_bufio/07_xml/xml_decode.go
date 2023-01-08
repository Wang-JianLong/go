package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

func main() {
	fileName := "Person.xml"
	file,_ := os.Open(fileName)
    decode := xml.NewDecoder(file)
	fmt.Println(decode.Token())
}