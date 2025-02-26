package leetcode263

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	for n%3 == 0 {
		n /= 3
	}
	for n%5 == 0 {
		n /= 5
	}
	return n&(n-1) == 0
}

// 首先，如果 n≤0，不符合题目正整数的要求，返回 false。
// 去掉 n 中的因子 3，也就是不断把 n 除以 3，直到 n 不是 3 的倍数为止。
// 去掉 n 中的因子 5，也就是不断把 n 除以 5，直到 n 不是 5 的倍数为止。
// 最后，n 必须只剩下因子 2，即 n=2
// k
//  。用 231. 2 的幂 中的技巧判断，见 我的题解。

// 作者：灵茶山艾府
// 链接：https://leetcode.cn/problems/ugly-number/solutions/3007807/li-yong-wei-yun-suan-you-hua-pythonjavac-nlqr/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
