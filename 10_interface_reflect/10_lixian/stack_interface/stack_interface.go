package main

import (
	"errors"
	"fmt"
)

type Top interface {
	Top() (interface{}, error)
}

const (
	EMPTY = 0
)

type IStack interface {
	Len() int
	IsEmpty() bool
	Push(x interface{})
	Pop() (x interface{}, e error)
}

type Stack []interface{}

func (s Stack) Len() int {
	return len(s)
}

func (s Stack) IsEmpty() bool {
	return s.Len() == EMPTY
}

func (s *Stack) Push(x interface{}) {
	*s = append(*s, x)
}

func (s Stack) Top() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("EMPTY!!!")
	}
	return s[s.Len()-1], nil
}

func (s *Stack) Pop() (x interface{}, e error) {
	stk := *s
	if stk.IsEmpty() {
		return nil, errors.New("EMPTY!!!")
	}
	ele, _ := s.Top()
	*s = stk[:s.Len()-1]
	return ele, nil
}

func main() {
	var s IStack = &Stack{}
	s.Push("A")
	s.Push(23)
	s.Push("asdasd")

	fmt.Printf("s: %v\n", s)
}
