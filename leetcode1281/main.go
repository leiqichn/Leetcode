/*
 * Copyright (c) 2025 Lei Qi. All rights reserved.
 * Author: Lei Qi
 * Description:
 * Date: 2025/2/23 上午11:11
 */

package leetcode1281

import "strconv"

func subtractProductAndSum(n int) int {
	nStr := strconv.Itoa(n)
	sums := 0
	products := 1

	for _, runeNum := range nStr {
		num := int(runeNum - '0')
		sums += num
		products *= num
	}
	return products - sums
}
