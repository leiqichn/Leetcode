package leetcode1_2

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		// 先检查再存入，避免自引用
		if idx, ok := m[complement]; ok {
			return []int{idx, i}
		}
		m[num] = i
	}
	return []int{}
}
