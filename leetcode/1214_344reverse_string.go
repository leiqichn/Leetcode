package leetcode

func reverseString(s []byte) {
	lens := len(s)
	for i, j := 0, lens-1; i < j; i++ {
		if i <= j {
			tmp := s[i]
			s[i] = s[j]
			s[j] = tmp
			j++
		}
	}
}
