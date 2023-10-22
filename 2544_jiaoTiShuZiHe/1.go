package _544_jiaoTiShuZiHe

import "strconv"

// 思路
// 找到最大值的index，然后正向和反向遍历
func alternateDigitSum(n int) int {
	nStr := strconv.Itoa(n)
	max := 0
	maxIdx := -1
	sum := 0
	for idx, ch := range nStr {
		ch2Int := int(ch - '0')
		if ch2Int > max {
			max = ch2Int
			maxIdx = idx
		}
	}

	for i := maxIdx; i < len(nStr); i++ {
		tmp := i - maxIdx
		if tmp%2 == 0 {
			sum += int(nStr[i] - '0')
		} else {
			sum -= int(nStr[i] - '0')
		}
	}
	for i := maxIdx; i >= 0; i-- {
		tmp := maxIdx - i
		if tmp%2 == 0 {
			sum += int(nStr[i] - '0')
		} else {
			sum -= int(nStr[i] - '0')
		}
	}
	sum -= int(nStr[maxIdx] - '0')
	return sum
}
