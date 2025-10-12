/*
 * Copyright (c) 2024 Lei Qi. All rights reserved.
 * Author: Lei Qi
 * Description:
 * Date: 2024/6/3 下午11:41
 */

package leetcode215

import "container/heap"

func findKthLargest(nums []int, k int) int {
	h := &IntHeap{}
	heap.Init(h)
	for _, v := range nums {
		heap.Push(h, v)
	}

	var res int

	for i := 0; i < k; i++ {
		res = heap.Pop(h).(int)
	}
	return res
}

// Len 给长度，Less 定大小，Swap 来交换，Push 加末尾，Pop 取最后
type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(*h)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
