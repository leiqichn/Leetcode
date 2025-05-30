package leetcode3335

func lengthAfterTransformations(s string, t int) int {
	// 一直进行一个操作 递归操作， 次数
	res := 0
	var dfs func([]byte, int) int
	dfs = func(s []byte, count int) int {
		if t == count {
			return len(s)
		}

		tmp_str := []byte{}

		for _, val := range s {
			if val == 'z' {
				tmp_str = append(tmp_str, 'a', 'b')
			} else {
				tmp_str = append(tmp_str, val+1)
			}
		}
		count++
		s = tmp_str

		return dfs(tmp_str, count)
	}
	strList := []byte(s)
	res = dfs(strList, 0)
	return res
}
