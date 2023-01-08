package main

// 下面的 Compare 函数会返回两个字节数组字典顺序的整数对比结果，
// 即 0 if a == b, -1 if a < b, 1 if a > b
func main() {

}

func Compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	// 数组的长度可能不同
	switch {
	case len(a) < len(b):
		return -1
	case len(a) > len(b):
		return 1
	}
	return 0 // 数组相等
}
