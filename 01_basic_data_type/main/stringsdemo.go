package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "abc"
	fmt.Println(strings.HasPrefix(str,"a"))
	fmt.Println(strings.HasSuffix(str,"c"))
	fmt.Println(strings.Replace(str,"b","AAA",1))

	s2 :="2"
	var s2n,_ = strconv.Atoi(s2);
	fmt.Println(s2n)
}