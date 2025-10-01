package leetcode

import "sort"

func topKFrequent(nums []int, k int) []int {
	keyValueMap := map[int]int{}
	ans := []int{}

	for _, k := range nums {
		if _, ok := keyValueMap[k]; ok {
			keyValueMap[k] += 1
		} else {
			keyValueMap[k] = 1
		}

	}
	for key, _ := range keyValueMap {
		ans = append(ans, key)
	}
	sort.Slice(ans, func(i, j int) bool {
		return keyValueMap[ans[i]] > keyValueMap[ans[j]]
	})

	return ans[:k]
}
