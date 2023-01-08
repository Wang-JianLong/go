package inter

import "fmt"

type OK interface {
	OK() OK
	Error() error
}

type DefualtOK struct {
	code int
	msg  string
}

func TestOk() *DefualtOK {
	return &DefualtOK{200, "SUCCESS"}
}

type Inner interface {
	Move()
}

type InnerImpl struct {
	w, h int
}

func (s *InnerImpl) Move() {
	fmt.Printf("\"我移动了一下\": %v\n", "我移动了一下")
}
