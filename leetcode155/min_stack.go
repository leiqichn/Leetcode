/*
 * Copyright (c) 2023-2024 Lei Qi. All rights reserved.
 * Author: Lei Qi
 * Description:
 * Date: 2024/5/8 上午12:55
 */

package leetcode155

import "math"

// 使用额外栈辅助， a b c d 入栈，如果d 在那么abc 必然在，所以每次入栈比较一个最小值，放入另一个栈就好
type MinStack struct {
	stack    []int // 主栈
	minStack []int // 辅助栈
}

func Constructor() MinStack {
	minStacks := MinStack{[]int{}, []int{math.MaxInt}}
	return minStacks
}

//func (this *MinStack) Push(x int) {
//	// 同时比较辅助站的的top元素，如果更小 则添加到辅助站
//	this.stack = append(this.stack, x)
//	minTmp := min(x, this.minStack[len(this.minStack)-1])
//	this.minStack = append(this.minStack, minTmp)
//}

func (this *MinStack) Push(x int) {
	// 同时比较辅助站的的top元素，如果更小 则添加到辅助站
	this.stack = append(this.stack, x)
	top := this.minStack[len(this.minStack)-1] // 需要使用top 暂存
	minTmp := min(x, top)
	this.minStack = append(this.minStack, minTmp)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return (this.stack[len(this.stack)-1])
}

func (this *MinStack) GetMin() int {
	return (this.minStack[len(this.minStack)-1])
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
