package stackstruct

import (
	"fmt"
	"strings"
)

type stack struct {
	nums [4]int "底层存放数据数组"
	flag int    "下标索引位置 不会大于 满了显示为4,空了显示为0"
}

func (s *stack) Push(data int) bool {
	if (s.flag + 1) > len(s.nums) {
		return false
	}
	s.nums[s.flag] = data
	s.flag += 1
	return true
}

func (s *stack) Pop() (bool, int) {
	num := 0
	if s.flag-1 < 0 {
		return false, num
	}
	s.flag -= 1
	num, s.nums[s.flag] = s.nums[s.flag], num
	return true, num
}

func (s *stack) String()string{
	var sb strings.Builder;
	for i,v := range s.nums{
		fmt.Printf(" [ %v : %v ] ",i,v)
	}
	return sb.String()
}

func NewStack() *stack{
	return &stack{}
}