package main

func main() {
	
	i := 0
STACKER_WJL:
	i++
	if i < 15 {
		println(i)
		goto STACKER_WJL
	}

}
