package leetcode2610

func findMatrix(nums []int) (ans [][]int) {
	cnt := make([]int, len(nums)+1)
	for _, x := range nums {
		c := cnt[x]
		if c == len(ans) { // 需要加一行
			ans = append(ans, []int{})
		}
		// 出现相同的则放到第几行
		ans[c] = append(ans[c], x)
		cnt[x]++
	}
	return
}
