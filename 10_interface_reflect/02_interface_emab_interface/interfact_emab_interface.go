package main

// 一个接口可以包含一个或多个其他的接口，这相当于直接将这些内嵌接口的方法列举在外层接口中一样
type Lock interface {
    Lock()
    Unlock()
}

type File interface {
    Lock
    Close()
}


func main(){
	var f File
	f.Lock()
	f.Close()
	f.Unlock()
}