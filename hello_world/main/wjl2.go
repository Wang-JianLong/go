package main

import (
	"errors"
	"fmt"
	"strings"
)

func GetErrorString() string {
	s := errors.New("这是一个一场 非常异常")
	fmt.Printf("s.Error(): %v\n", s.Error())

return	strings.Join(strings.Split(s.Error(), "asassdad"), "sdadasd")
}
