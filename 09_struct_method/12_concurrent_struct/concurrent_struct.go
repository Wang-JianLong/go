package main

import (
	"fmt"
	"sync"
)

func main() {
	p :=new(Person)
	p.mu.Lock()
	p.SetFirstName("张三")
	fmt.Printf("p: %v\n", p)
	defer p.mu.Unlock()
}

type Person struct {
	firstName, lastName string
	mu                  sync.Mutex // 初始化时不需要赋值
}

// 对于 setter 方法使用 Set 前缀，对于 getter 方法只使用成员名。
func (self *Person) SetFirstName(firstName string) {
	self.firstName = firstName
}

func (self *Person) FirstName() string {
	return self.firstName
}
