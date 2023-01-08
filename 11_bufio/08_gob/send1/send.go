package main

import (
	"encoding/gob"
	"os"
)

func main() {
	var t = T{X: 7, Y: 0, Z: 8}
	out,_ := os.OpenFile("gob.txt",os.O_WRONLY|os.O_CREATE,0666)
	
	encode := gob.NewEncoder(out)
	encode.Encode(t)
}

type T struct {
	X, Y, Z int
}
