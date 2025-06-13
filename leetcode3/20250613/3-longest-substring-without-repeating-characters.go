package _0250613

func lengthOfLongestSubstring(s string) int {

	maps := map[byte]int{}

	left := 0
	res := 0
	for right, b := range s {
		maps[byte(b)] += 1
		for _, c := range maps {
			for c > 1 { ///不用判断所有map 中 滑动窗口的单调性保证
				//			滑动窗口的核心思想是：当窗口满足条件时（无重复字符），右指针向右扩展；当窗口不满足条件时（出现重复），左指针向右收缩。
				//
				//			关键观察：如果当前字符 b 的计数变为 2，说明它一定是刚刚被右指针引入的字符，且之前已经存在于窗口中。此时：
				//
				//			这个重复的 b 一定是窗口中最右侧的重复字符（因为右指针刚加入它）。
				//
				//			窗口中其他字符的计数最多为 1（否则在之前的迭代中左指针已经移动了）。
				//
				//			因此，只需要检查当前字符 b 是否重复，就能确定是否需要收缩窗口。
				left++
				maps[s[left]]--
			}

		}
		res = max(res, right-left+1)
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

//func lengthOfLongestSubstring(s string) int {
//	maps := map[rune]int{}
//
//	left := 0
//	res := 0
//	for right, b := range s {
//		maps[b] += 1
//		for _, c := range maps { //判断所有map 中 滑动窗口的单调性保证
//			滑动窗口的核心思想是：当窗口满足条件时（无重复字符），右指针向右扩展；当窗口不满足条件时（出现重复），左指针向右收缩。
//
//			关键观察：如果当前字符 b 的计数变为 2，说明它一定是刚刚被右指针引入的字符，且之前已经存在于窗口中。此时：
//
//			这个重复的 b 一定是窗口中最右侧的重复字符（因为右指针刚加入它）。
//
//			窗口中其他字符的计数最多为 1（否则在之前的迭代中左指针已经移动了）。
//
//			因此，只需要检查当前字符 b 是否重复，就能确定是否需要收缩窗口。
//			for c > 1 {
//				maps[rune(s[left])]--
//				left++
//			}
//			res = max(res, right-left+1)
//		}
//
//	}
//	return res
//}
//
//func max(a, b int) int {
//	if a < b {
//		return b
//	}
//	return a
//}
