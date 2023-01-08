package pack1

import "fmt"

var (
	Pack1Int   int     = 1
	Pack1Float float64 = 1.1
)

func ReturnStr() string {
	return "Hello main!"
}

func init() {
	fmt.Printf("\"pack1 Init...\": %v\n", "pack1 Init...")
}