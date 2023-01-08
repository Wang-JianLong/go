// 切片重组
package main

/*
	我们已经知道切片创建的时候通常比相关数组小
	slice1 := make([]type, start_length, capacity)
	其中 start_length 作为切片初始长度而 capacity 作为相关数组的长度。

这么做的好处是我们的切片在达到容量上限后可以扩容。改变切片长度的过程称之为切片重组 reslicing，做法如下：slice1 = slice1[0:end]，其中 end 是新的末尾索引（即长度）
*/
func main() {
	// 将切片扩展 1 位可以这么做
	sl := make([]int, 10, 20)
	// sl = sl[0 : len(sl)+1] // 切片可以反复扩展直到占据整个相关数组。
	
}
