package main

import (
	"fmt"
	"os"
)

func main() {
		
}

func cat(f *os.File) {
	const NBUF = 512
	var buf [NBUF]byte

	for {
		switch nr, err := f.Read(buf[:]); true {
			case nr < 0:
				fmt.Fprintf(os.Stderr, "cat: error read: %s\n", err.Error())
				os.Exit(1)
			case nr == 0: // EOF
				return
			case nr > 0:
				if nw, ew := os.Stdout.Write(buf[0:nr]); nw != nr {
					fmt.Fprint(os.Stderr, "cat: err writing:%s\n", ew.Error())
				}
			}
	}
}
