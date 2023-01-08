package main

import (
	"encoding/csv"
	"fmt"
	"strings"
)

func main() {
	msg := `font;fontSize;
fira Code;h12;`
	read := csv.NewReader(strings.NewReader(msg))

	for{
		record,err := read.Read()
		if err != nil{
			break
		}
		fmt.Printf("record: %v\n", record)
	}
}