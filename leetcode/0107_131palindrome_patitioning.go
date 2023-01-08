package leetcode

// 给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
//
// 回文串 是正着读和反着读都一样的字符串。
//
// 示例 1：
//
// 输入：s = "aab"
// 输出：[["a","a","b"],["aa","b"]]
// 示例 2：
//
// 输入：s = "a"
// 输出：[["a"]]
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/palindrome-partitioning
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
var (
	path []string
	res  [][]string
)

func partition(s string) [][]string {
	path, res = make([]string, 0), make([][]string, 0)
	backTracking(s, 0)
	return res
}

// 弄清楚输入输出 需要
func backTracking(s string, startIndex int) {
	// 终止条件
	if startIndex >= len(s) { // 需要减1吗
		var tmp []string
		tmp = make([]string, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		return
	}

	for i := startIndex; i < len(s); i++ {
		// 单个循环的操作这里是横向的操作 递归是纵向的
		// 子串为【startIndex，i】
		// 判断是否是回文子串
		subStr := s[startIndex : i+1]
		if isReveredStr(subStr) {
			// 添加 path
			path = append(path, subStr)
			backTracking(s, i+1) // 这里为什么不使用 startIndex +1呢？
			path = path[:len(path)-1]
		} else {
			continue
		}

	}

}

// 双指针判断是否是回文字串
// str 可以直接使用 切片
func isReveredStr(str string) bool {
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 { // 只能这样多个赋值
		if str[i] != str[j] {
			return false
		}
	}
	return true
}
