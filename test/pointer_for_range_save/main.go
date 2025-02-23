/*
 * Copyright (c) 2024-2025 Lei Qi. All rights reserved.
 * Author: Lei Qi
 * Description:
 * Date: 2025/2/23 上午11:10
 */

package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {
	students := []*Student{
		{Name: "Alice", Age: 20},
		{Name: "Bob", Age: 22},
	}

	// 使用for range遍历指针数组
	for _, student := range students {
		// 修改指向的结构体的内容
		student.Age += 1 // 给每个学生的年龄加1
	}

	// 打印修改后的students数组
	for _, student := range students {
		fmt.Printf("Name: %s, Age: %d\n", student.Name, student.Age)
	}
}
