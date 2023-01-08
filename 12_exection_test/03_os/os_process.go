package main

import (
	"fmt"
	"os"
)

func main(){
	env := os.Environ()
	attr := &os.ProcAttr{
		Env:env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}

	pid,err :=	os.StartProcess("/bin/ls",[]string{"ls","-l"},attr)
	if err != nil {
		fmt.Printf("Error %v",err)
	}

	fmt.Printf("pid: %v\n", pid)
	
}

