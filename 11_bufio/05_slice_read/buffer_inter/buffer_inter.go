package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Fprint(os.Stdout,"123")
	
	bu := bufio.NewWriter(os.Stdout)
	fmt.Fprint(bu,1,2,3)
	bu.Flush()
	
}