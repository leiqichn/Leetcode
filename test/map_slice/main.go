/*
 * Copyright (c) 2024 Lei Qi. All rights reserved.
 * Author: Lei Qi
 * Description:
 * Date: 2024/12/3 下午11:24
 */
package main

import "fmt"

func main() {
	// 定义一个map，其key是一个长度为3的int数组
	var m = make(map[[3]int]string)

	// 向map中添加元素
	m[[3]int{1, 2, 3}] = "one-two-three"
	m[[3]int{4, 5, 6}] = "four-five-six"

	// 从map中读取元素
	fmt.Println(m[[3]int{1, 2, 3}]) // 输出: one-two-three
	fmt.Println(m[[3]int{4, 5, 6}]) // 输出: four-five-six

	// 检查一个不存在的key
	if val, ok := m[[3]int{7, 8, 9}]; ok {
		fmt.Println(val)
	} else {
		fmt.Println("Key not found")
	}
}
