package facotry

/*
Go 语言不支持面向对象编程语言中那样的构造子方法，但是可以很容易的在 Go 中实现 “构造子工厂” 方法。
为了方便通常会为类型定义一个工厂，按惯例，工厂的名字以 new 或 New 开头。
*/

type File struct{
	fd int  // 文件描述符
	name string	
}

func NewFile(fd int,name string) *File{
	return &File{fd,name}
}

// func main(){
// 	// 在 Go 语言中常常像上面这样在工厂方法里使用初始化来简便的实现构造函数	
// 	f := NewFile(1,"1.txt")
// 	fmt.Printf("unsafe.Sizeof(f): %v\n", unsafe.Sizeof(f))

// 	max :=max{}

// }

// 为了保证工厂的私密性 禁止其他包使用new() 可以小写结构体
type max struct{

}

func NewMax() *max{
	return &max{}
}
