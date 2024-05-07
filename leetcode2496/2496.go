/*
 * Copyright (c) 2024 Lei Qi. All rights reserved.
 * Author: Lei Qi
 * Description:
 * Date: 2024/5/8 上午12:55
 */

package leetcode2496

import (
	"fmt"
	"strconv"
)

func maximumValue(strs []string) int {
	// 只包含数字
	max := 0
	for i := 0; i < len(strs); i++ {
		if isNum(strs[i]) {
			str2int, err := strconv.Atoi(strs[i])
			if err != nil {
				fmt.Println("error")
			}
			if max < str2int {
				max = str2int
			}
		} else {
			lens := len(strs[i])
			if max < lens {
				max = lens
			}
		}
	}
	return max
}

func isNum(strs string) bool {
	for _, i := range strs {
		if i > '9' || i < '0' {
			return false
		}
	}
	return true
}
