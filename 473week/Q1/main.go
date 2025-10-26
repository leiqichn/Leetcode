package Q1

import "strconv"

func removeZeros(n int64) int64 {
	// 移除所有0
	// 转化为字符串， 双指针移除所有0， 拼接

	strs := strconv.Itoa(int(n))

	tmps := []byte(strs)

	res := []byte{}

	for i := 0; i < len(tmps); i++ {
		if tmps[i] == '0' {
			continue
		}
		res = append(res, tmps[i])
	}
	resStr := string(res)

	resInt, _ := strconv.Atoi(resStr)

	return int64(resInt)

}
