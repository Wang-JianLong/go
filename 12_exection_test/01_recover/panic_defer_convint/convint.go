package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(IntFrmint64(math.MaxInt64))
}

func ConverInt64ToInt(x int64) int {
	if x <= math.MaxInt32 && x >= math.MinInt32 {
		return int(x)
	}
	panic(fmt.Sprintf("%d is out of the interge32", x))
}

func IntFrmint64(x int64) (s int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	s = ConverInt64ToInt(x)
	return
}
