/*
 * Copyright (c) 2024 Lei Qi. All rights reserved.
 * Author: Lei Qi
 * Description:
 * Date: 2024/6/3 下午11:41
 */

package leetcode215

import "container/heap"

func findKthLargest(nums []int, k int) int {
	h := heapify(nums)
	var res any
	for i := 0; i < k; i++ {
		res = heap.Pop(&h)
	}
	return res.(int)
}

type BigHeap []int

func (h BigHeap) Len() int { return len(h) }
func (h BigHeap) Less(i, j int) bool {
	// 大根堆
	return h[i] > h[j]
}

func (h BigHeap) Swap(i, j int) {
	tmp := h[i]
	h[i] = h[j]
	h[j] = tmp
}
func (h *BigHeap) Push(x any) {
	*h = append(*h, x.(int))
}

// 删除元素待定
func (h *BigHeap) Pop() any {
	x := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return x
}

// 将 nums 转换成 BigHeap
func heapify(nums []int) BigHeap {
	h := make(BigHeap, len(nums)) // 新建BigHeap，长度为lenNums
	copy(h, nums)                 // 将num copy 到 BigHeap 中去
	heap.Init(&h)                 // 需要输入指针
	return h
}
