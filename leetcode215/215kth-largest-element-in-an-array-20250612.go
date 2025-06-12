/*
 * Copyright (c) 2024 Lei Qi. All rights reserved.
 * Author: Lei Qi
 * Description:
 * Date: 2024/6/3 下午11:41
 */

package leetcode215

import (
	"container/heap"
)

func findKthLargest(nums []int, k int) int {
	h := IntHeap{} // 可以直接实例化
	heap.Init(&h)  // 需要输入指针 初始化之后，可使用push
	for _, v := range nums {
		heap.Push(&h, v) // 使用指针
	}
	var res int
	for i := 0; i < k; i++ {
		res = heap.Pop(&h).(int)
	}
	return res
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h *IntHeap) Push(i any) {
	*h = append(*h, i.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
