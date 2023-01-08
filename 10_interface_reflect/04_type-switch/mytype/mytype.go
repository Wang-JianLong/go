package main

import "fmt"

type IService interface {
	Appect()
}

type Service struct {
}

func (s *Service) Appect() {

}
func (s *Service) Self() {
	fmt.Printf("\"我自己\": %v\n", "我自己")
}

func main() {
	var s IService = &Service{}
	switch t := s.(type) {
	case *Service:
		t.Self()
	}

	a := s.(*Service)
	a.Self()
	fmt.Printf("a: %v\n", a)
}