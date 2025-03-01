package leetcode1422

func maxScore(s string) int {
	// 最简单的方法 暴力遍历 res 更新
	// 前缀和
	res := 0
	zeroNumber := 0
	OneNumber := 0

	zeroNumberList := make([]int, len(s)-1)
	OneNumberList := make([]int, len(s)+1)
	// // 前缀和 注意是非空字符串，字符串长度只能取n-1 len(s)-1 要留最后一个
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '0' {
			zeroNumber += 1
			zeroNumberList[i] = zeroNumber
		}
	}

	// 前缀和 注意是非空字符串，字符串长度只能取n-1 for i:=0; i <len(s)-1; i++    for i:=len(s)-1; i > 0(不能等于0);
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == '1' {
			OneNumber += 1
			OneNumberList[i-1] = OneNumber // 错开，需要将i+1的位置值放到i 上
		}
	}

	for i := 0; i < len(s)-1; i++ {
		tmp := zeroNumberList[i] + OneNumberList[i]
		if tmp > res {
			res = tmp
		}
	}
	return res
}
