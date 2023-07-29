package leetcode

func numJewelsInStones(jewels string, stones string) int {
	stoneMaps := make(map[rune]int, len(stones))
	count := 0
	for _, val := range stones {
		stoneMaps[val] += 1
	}

	for _, stone := range jewels {
		if val, ok := stoneMaps[stone]; ok {
			count += val
		}
	}

	return count
}
