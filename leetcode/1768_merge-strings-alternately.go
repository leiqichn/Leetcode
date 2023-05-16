package leetcode

func mergeAlternately(word1 string, word2 string) string {
	str, lens := lensMax(word1, word2)
	res := []byte{}
	if str == word1 {
		// 遍历短的word2
		// 多余的部分
		last := word1[len(word1)-lens:]
		for i := 0; i < len(word2); i++ {

			res = append(res, word1[i])
			res = append(res, word2[i])
		}
		res = append(res, last...)
	} else if str == word2 {
		last := word2[len(word2)-lens:]
		for i := 0; i < len(word1); i++ {

			res = append(res, word1[i])
			res = append(res, word2[i])
		}
		res = append(res, last...)
	}
	return string(res)
}

// 判断谁的长度长，返回长的那个字符串
func lensMax(word1 string, word2 string) (str string, lens int) {
	if len(word1) > len(word2) {
		lens = len(word1) - len(word2)
		return word1, lens
	}
	lens = len(word2) - len(word1)
	return word2, lens
}
