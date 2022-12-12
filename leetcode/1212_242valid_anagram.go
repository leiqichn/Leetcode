package leetcode

import (
	"sort"
	"strings"
)

func isAnagram(s string, t string) bool {
	// 次数相等 使用map 存储次数
	//sort.Slice(s ,func(i,j int) bool{
	//
	//})
	if len(s) != len(t) {
		return false
	}
	//对字符串排序的方法
	//1.将字符串进行切片，转换为切片
	sSplit := strings.Split(s, "")
	//2.对切片进行排序（使用自带的排序库sort.Strings([] string{})即可）
	sort.Strings(sSplit)
	tSplit := strings.Split(t, "")
	sort.Strings(tSplit)
	//3.使用将切片 数组不能直接比较，需要逐个元素一次比较，需要使用双指针，终止条件是遍历结束，然后比较每个元素，之后两个指针均加1
	//strings.Join([]string,"")
	//p1, p2 := 0, 0
	//for p1 < len(s) {
	//	if string(sSplit[p1]) != string(tSplit[p2]) {
	//		return false
	//	}
	//	p1++
	//	p2++
	//}
	//3 或者再将其拼接为字符串直接进行等于比较
	sJoin := strings.Join(sSplit, "")
	tJoin := strings.Join(tSplit, "")
	if sJoin != tJoin {
		return false
	}
	return true
}

func isAnagramByMap(s string, t string) bool {
	chrMap := make(map[rune]int, len(s))

	for _, val := range s {
		if _, ok := chrMap[val]; ok {
			chrMap[val]++
		} else {
			chrMap[val] = 1
		}
	}
	for _, val := range t {
		if _, ok := chrMap[val]; ok {
			chrMap[val]--
		} else {
			return false
		}
	}

	for _, count := range chrMap {
		if count != 0 {
			return false
		}
	}
	return true

}
