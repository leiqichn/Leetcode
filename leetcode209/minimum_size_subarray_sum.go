package leetcode

import "math"

//给定一个含有 n 个正整数的数组和一个正整数 target 。
//
// 找出该数组中满足其总和大于等于 target 的长度最小的 子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其
//长度。如果不存在符合条件的子数组，返回 0 。
//
//
//
// 示例 1：
//
//
//输入：target = 7, nums = [2,3,1,2,4,3]
//输出：2
//解释：子数组 [4,3] 是该条件下的长度最小的子数组。
//
//
// 示例 2：
//
//
//输入：target = 4, nums = [1,4,4]
//输出：1
//
//
// 示例 3：
//
//
//输入：target = 11, nums = [1,1,1,1,1,1,1,1]
//输出：0
//
//
//
//
// 提示：
//
//
// 1 <= target <= 10⁹
// 1 <= nums.length <= 10⁵
// 1 <= nums[i] <= 10⁵
//
//
//
//
// 进阶：
//
//
// 如果你已经实现 O(n) 时间复杂度的解法, 请尝试设计一个 O(n log(n)) 时间复杂度的解法。
//
//
// Related Topics 数组 二分查找 前缀和 滑动窗口 👍 2210 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
// 暴力解法
func minSubArrayLen1(target int, nums []int) int {
	minLen := math.MaxInt
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum >= target {
				minLen = min(minLen, j-i+1)
				break
			}
		}
	}
	if minLen == math.MaxInt {
		return 0
	}
	return minLen
}

func minSubArrayLen(target int, nums []int) int {
	// 暴力解法 是两层遍历，每次更新最小的res
	// 滑动窗口，则需要 先遍历 确定终止条件，然后再while 不断更新起始位置，把最小的结果更新到res 里边
	minRes := 1<<63 - 1
	i := 0
	sum := 0
	for j := 0; j < len(nums); j++ { // 先确定终点
		sum += nums[j]
		for sum >= target { // 一直循环sum >= target 这个条件
			sublength := j - i + 1  // 求的是长度，需要计算长度
			if sublength < minRes { // 求最小长度
				minRes = sublength
			}
			sum -= nums[i] // 需要不断缩小i,sum 就需要不断减去nums[i]
			i++            // 在符合的终点的循环里边，不断缩小起点
		}
	}
	if minRes == 1<<63-1 {
		return 0
	}
	return minRes

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
