package leetcode

//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//
//
//
// 示例 1：
//
//
//
//
//输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
//输出：6
//解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
//
//
// 示例 2：
//
//
//输入：height = [4,2,0,3,2,5]
//输出：9
//
//
//
//
// 提示：
//
//
// n == height.length
// 1 <= n <= 2 * 10⁴
// 0 <= height[i] <= 10⁵
//
//
// Related Topics 栈 数组 双指针 动态规划 单调栈 👍 5266 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func trap(height []int) int {
	left := 0
	right := len(height) - 1
	pre_max := 0
	suf_max := 0
	res := 0
	for left <= right {
		pre_max = max(pre_max, height[left])
		suf_max = max(suf_max, height[right])
		if pre_max > suf_max {
			res += suf_max - height[right]
			right -= 1
		} else {
			res += pre_max - height[left]
			left += 1
		}
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
