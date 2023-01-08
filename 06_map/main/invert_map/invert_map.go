package main

import (
	"fmt"
	"sort"
)

/*
这里对调是指调换 key 和 value。如果 map
的值类型可以作为 key 且所有的 value 是唯一的，那么通过下面的方法可以简单的做到键值对调。
*/
var (
    barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
                            "delta": 87, "echo": 56, "foxtrot": 12,
                            "golf": 34, "hotel": 16, "indio": 87,
                            "juliet": 65, "kili": 43, "lima": 98}
)


func main(){
	invMap := make(map[int]string,len(barVal))

	for k,v := range barVal {
		invMap[v] = k
	}

	fmt.Printf("invMap: %v\n", invMap)
	// invMap: map
	// [12:foxtrot 16:hotel 23:charlie 34:alpha 43:kili 56:bravo 65:juliet 87:indio 98:lima]

	/*
	如果原始 value 值不唯一那么这么做肯定会出错；
		为了保证不出错，当遇到不唯一的 key 时应当立刻停止，
		这样可能会导致没有包含原 map 的所有键值对！
		一种解决方法就是仔细检查唯一性并且使用多值 map，
		比如使用 map[int][]string 类型
	*/

	PrintDrink()
}

/*构造一个将英文饮料名映射为法语（或者任意你的母语）的集合；
	先打印所有的饮料，
	然后打印原名和翻译后的名字。
	接下来按照英文名排序后再打印出来。
*/
func PrintDrink(){
	drinks := map[string]string{
		"coffee":"咖啡",
		"tea":"茶",
		"milk tea":"奶茶",
		"milk":"牛奶",
	}

	for _,v := range drinks {
		fmt.Printf("v: %v\n", v)
	}

	for k,v := range drinks{
		fmt.Printf("原名: %v	翻译: %v \n",v,k)
	}

	keys := make([]string,len(drinks))
	i := 0
	for k := range drinks{
		keys[i] = k
		i++
	}

	sort.Strings(keys)
	fmt.Printf("keys: %v\n", keys)

}