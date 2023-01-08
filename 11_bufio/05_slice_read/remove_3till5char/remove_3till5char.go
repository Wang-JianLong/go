package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, _ := os.Open("goprogram")
	outputFile, _ := os.OpenFile("goprogramT", os.O_WRONLY|os.O_CREATE, 0666)
	defer inputFile.Close()
	defer outputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	outputWriter := bufio.NewWriter(outputFile)
	
	for {
		inputReader.ReadLine()
	}
	
	fmt.Println("Conversion done")
}
