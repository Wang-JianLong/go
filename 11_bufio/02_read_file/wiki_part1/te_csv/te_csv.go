package main

import (
	"encoding/csv"
	"os"
)

func main() {
	out,_:= os.OpenFile("aaa.abc",os.O_WRONLY|os.O_CREATE,0666)
	csvW := csv.NewWriter(out)
	msg := []string{"state","msg","1","2"}
	csvW.Write(msg)	
	csvW.Flush()
}