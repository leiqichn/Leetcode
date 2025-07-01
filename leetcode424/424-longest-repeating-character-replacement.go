package leetcode424

func characterReplacement(s string, k int) int {
	// 相同字母判断
	keyMap := make([]int, 26)
	res := 0
	left := 0
	maxCount := 0
	for right, ch := range s {
		key := ch - 'A'
		keyMap[key]++
		if keyMap[key] > maxCount {
			maxCount = keyMap[key]
		}

		if (right-left+1)-maxCount > k {
			keyMap[s[left]-'A']--
			left++
		}

		winSize := right - left + 1

		res = max(winSize, res)

	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
