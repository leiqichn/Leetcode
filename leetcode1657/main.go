package leetcode1657

func closeStrings(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var sCnt, tCnt [26]int
	for _, c := range s {
		sCnt[c-'a']++
	}
	for _, c := range t {
		tCnt[c-'a']++
	}

	for i := range 26 {
		// 判断 s 和 t 的字符集合是否一样，如果不一样直接返回 false。例如 s 中有字符 abc，t 中有字符 def，我们无论如何都不能把 s 变成 t
		// 判断对应字符数量为0， 对应自符数量不为0， 巧妙
		if (sCnt[i] == 0) != (tCnt[i] == 0) {
			return false
		}
	}

	slices.Sort(sCnt[:]) // 排序集合数量
	slices.Sort(tCnt[:]) // 排序结合数量， 可以交换数量
	return sCnt == tCnt
}
