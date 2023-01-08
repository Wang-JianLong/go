package inter

type Sorter interface {
	Less(i, j int) bool //  判断后一位i是否 是否大于前一位 j
	Swap(i, j int)      // 交换两个位置的值
	Len() int           // 要排序的实现类的长度
}

func SortFun(s Sorter) {
	for l := 1; l < s.Len(); l++ {
		for j := 0; j < s.Len()-l; j++ {
			if s.Less(j+1, j) {
				s.Swap(j, j+1)
			}
		}
	}
}

func IsSorted(s Sorter) bool {
	for i := 1; i < s.Len(); i++ {
		for j := 0; j < s.Len()-i; j++ {
			if !s.Less(j+1, j) {
				return false
			}
		}
	}
	return true
}
