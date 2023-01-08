package main
import "fmt"
func main() {
	i := []int{1,3,4,6,23,5,3,2,9,7,5}
	i = sort3(i)
	fmt.Printf("%v\n",i)
}

/*
	使用冒泡排序的方法排序一个包含整数的切片
*/

func sort(s []int) []int{
	for i := 0; i < len(s); i++ {
        for j := 1; j < len(s) - i; j++{
			if s[j] < s[j - 1] {
				s[j-1],s[j] = s[j],s[j-1]
			} 
		}
	}
	return s
}

func sort2(s []int) []int{
	for i := 0; i < len(s) ; i++ {
		swaped := false
        for j := 0; j < len(s) - i - 1; j++{
			if s[j] > s[j + 1] { 
				s[j+1],s[j] = s[j],s[j+1]
				swaped = true
			} 
		}
		if (!swaped) {
			break;
		}
	}
	return s
}

func sort3(s []int) []int{
	n := len(s) - 1
	for {
		last := 0
		for i := 0; i < n;i++{
			if s[i] > s[i + 1] { 
				s[i+1],s[i] = s[i],s[i+1]
				last = i
			} 
		}
		if last == 0 {
			break
		}
		n = last
	}
	return s
}

func bubblesort(a []int) []int {
	for i := 1; i < len(a); i++ {
		for j := 1; j < len(a)-i+1; j++ {
			if a[j] < a[j-1] {
				temp := a[j]
				a[j] = a[j-1]
				a[j-1] = temp
			}
		}
	}
	return a
}
