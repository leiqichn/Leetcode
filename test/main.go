/*
 * Copyright (c) 2024 Lei Qi. All rights reserved.
 * Author: Lei Qi
 * Description: 测试文件
 * Date: 2024/11/23 下午1:55
 */

package test

import (
	"fmt"
	"strings"
)

func main() {
	haystack := "good"
	needle := "go"
	splitList := strings.Split(haystack, needle) // 如果在首位相同，前边会有个空字符串。
	fmt.Println(splitList, len(splitList))
	fmt.Println("tmp:", splitList[0])
}
