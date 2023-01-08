package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func main() {
	var network bytes.Buffer
	// Buffer 即时读也是写
	enc := gob.NewEncoder(&network)
	dec := gob.NewDecoder(&network)

	err := enc.Encode(P{3, 4, 5, "JAVA"})
	if err != nil {
		// ...
	}

	var q Q;
	err = dec.Decode(&q)
	if err != nil{
		//...
	}

	fmt.Printf("q.Name: %v\n", q.Name) // 没有传输
	x := q.X
	print(*x)
	y := q.Y
	print(*y)
}


type P struct {
	X, Y, Z int
	Names   string
}

type Q struct {
	X, Y *int32
	Name string
}
