/*
 * Copyright (c) 2025 Lei Qi. All rights reserved.
 * Author: Lei Qi
 * Description:
 * Date: 2025/2/23 上午11:11
 */

package main

import (
	"fmt"
	"strconv"
)

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

func subtractProductAndSum2(n int) int {
	m := 1
	s := 0
	for n > 0 { // 当n /10 == 0 的时候可以返回了
		x := n % 10 // n%10 是获取个位上的数字
		n /= 10     // 除10 得到的商是相当于把各位的数给删掉了
		m *= x
		s += x
	}
	return m - s
}

// 递归解法

func getSums(n int) int {
	// base case 当只剩一位的时候，n/10 == 0
	if n/10 == 0 {
		return n
	}
	sums := (n % 10) + getSums(n/10) // n%10 是当前位数，接下来需要输入去掉最后一位
	return sums
}

func getProducts(n int) int {
	// base case
	if n/10 == 0 {
		return n
	}
	products := (n % 10) * getProducts(n/10)
	return products
}

func subtractProductAndSumRecursion(n int) int {
	return getProducts(n) - getSums(n)
}

func main() {
	subtractProductAndSum2(123)
	subtractProductAndSum(123)
	fmt.Println(1 / 10)
}
