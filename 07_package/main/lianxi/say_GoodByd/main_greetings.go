package main

import (
	"07_package/main/lianxi/say_GoodByd/greeting"
)
func main(){
	if greeting.IsAm() {
		greeting.SayEvening()
	}else{
		greeting.SayAfternoon()
	}
}