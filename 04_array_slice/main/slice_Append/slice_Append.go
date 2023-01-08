package main
import "fmt"
func main() {
	s := []byte{1,2,3}
	s = Append(s,4,5,6)
	fmt.Printf("cap(s): %v\n", cap(s))
	fmt.Printf("len(s): %v\n", len(s))
    fmt.Printf("%v\n",s)
}

func Append(sl []byte,data ...byte) []byte {
	m := len(sl) // 3
	// 计算出总和长度
	n := m + len(data) // 6
	// 比较大就扩容
    if n > cap(sl) {
		newSlice := make([]byte, (n+1) * 2) // cap 14 len 14
		// 将原切片copy至新数组
		copy(newSlice, sl)
        sl = newSlice
	}
	// 子切片 为之前的总长度
	sl = sl[:n] // len 6 cap 14
	// sl[3:6] 又是一个子切片 并且为母切片的后三位
	copy(sl[m:n],data) // 拷贝到这个切片的后三位 也就是 data的预备长度
	return sl // 注意 这里返回的是母切片
}
