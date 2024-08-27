/*
 * Copyright (c) 2024 Lei Qi. All rights reserved.
 * Author: Lei Qi
 * Description:
 * Date: 2024/8/27 下午11:26
 */

package leetcode3

func lengthOfLongestSubstring(s string) int {
	ans := 0
	cnt := make(map[rune]int, len(s))
	left := 0

	for right, c := range s {
		cnt[c] += 1      // 由于之前都是保证了不是重复的，所以只可能是右边新加的是重复的
		for cnt[c] > 1 { // 有重复，则不断的将左端点右移
			cnt[rune(s[left])] -= 1
			left++
		}
		// 没有重复的时候， 计算字符个数
		ans = max(ans, right-left+1)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
