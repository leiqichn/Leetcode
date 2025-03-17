package leetcode1963

func minSwaps(S string) (ans int) {
	s := []byte(S)
	c := 0
	j := len(s) - 1
	for _, b := range s {
		if b == '[' {
			c++
		} else if c > 0 {
			c--
		} else { // c == 0
			// 找最右边的左括号交换
			for s[j] == ']' {
				j--
			}
			s[j] = ']' // s[i] = '[' 可以省略
			ans++
			c++ // s[i] 变成左括号，c 加一
		}
	}
	return
}
