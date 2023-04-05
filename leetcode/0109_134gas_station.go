package leetcode

func canCompleteCircuit(gas []int, cost []int) int {
	// 暴力模拟

	for i := 0; i < len(gas); i++ {
		//模拟每个起点
		totalRest := 0
		rest := 0
		index := 0
		for rest > 0 && index != i {
			index = i
			rest = gas[i] - cost[i]
			index = (i + 1) % len(gas)
			totalRest += rest
		}
		if totalRest > 0 {
			return i
		}
	}
	return -1

}
