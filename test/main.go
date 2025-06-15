/*
 * Copyright (c) 2024 Lei Qi. All rights reserved.
 * Author: Lei Qi
 * Description: 测试文件
 * Date: 2024/11/23 下午1:55
 */

package main

import (
	"fmt"
	"strings"
)

func main() {
	haystack := "good"
	needle := "go"
	splitList := strings.Split(haystack, needle) // 如果在首位相同，前边会有个空字符串。
	fmt.Println(splitList, len(splitList))
	fmt.Println("tmp:", splitList[0], splitList[1])

	// 验证切片传参修改
	a := []int{1, 2, 3}
	modifySlice(a)
	fmt.Println(a) // 输出: [99 2 3]

	modifySliceByPtr(&a)
	fmt.Println(a) // 输出: [99 2 3 100]
}

func modifySlice(s []int) {
	s = append(s, 100) // 不会影响外部
	s[0] = 99          // 会影响外部（修改底层数组）
}

func modifySliceByPtr(s *[]int) {
	*s = append(*s, 100) // 会影响外部
}
