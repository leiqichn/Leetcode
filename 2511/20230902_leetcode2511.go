package _511

// 20230902
func captureForts(forts []int) int {
	// 思路 分别遍历 从1 找 -1 和 -1 开始 -1 到1 ，中间需要都是0
	//　最终结果　res int
	var res int = 0
	for i := 0; i < len(forts); i++ {
		if forts[i] == 1 {
			tmpZeroLen := 0
			for j := i + 1; j < len(forts); j++ {
				if forts[j] == 0 && (forts[j-1] == 0 || (j-1 == i)) {
					tmpZeroLen++
					continue
				} else if forts[j] == 0 || forts[j] == 1 {
					tmpZeroLen = 0
					break
				} else if forts[j] == -1 {
					res = max(res, tmpZeroLen)
					break
				}

			}
		} else if forts[i] == -1 {
			tmpZeroLenRev := 0
			for j := i + 1; j < len(forts); j++ {
				if forts[j] == 0 && (forts[j-1] == 0 || (j-1 == i)) {
					tmpZeroLenRev++
					continue
				} else if forts[j] == 0 || forts[j] == -1 {
					tmpZeroLenRev = 0
					break
				} else if forts[j] == 1 {
					res = max(res, tmpZeroLenRev)
					break
				}
			}
		}

	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
