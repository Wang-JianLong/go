package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	md := md5.New()
	io.WriteString(md,"Hash")
	b := []byte{}
	fmt.Printf("md.Sum(b): %v\n", md.Sum(b))
}
