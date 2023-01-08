package person

type Person struct{
	firstName,lastName string
}

// 对于 setter 方法使用 Set 前缀，对于 getter 方法只使用成员名。
func (self *Person) SetFirstName(firstName string){
	self.firstName = firstName
}

func (self *Person) FirstName() string{
	return self.firstName
}