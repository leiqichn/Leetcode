/*
 * Copyright (c) 2024 Lei Qi. All rights reserved.
 * Author: Lei Qi
 * Description:
 * Date: 2024/6/3 下午11:41
 */

package leetcode295

import "container/heap"

//import ""

type maxHeap []int

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	(*h) = (*h)[:n-1]
	return x
}

type minHeap []int

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := (*h)[n-1]
	(*h) = (*h)[:n-1]
	return x
}

// 中位数为
// num <= 中位数 maxQue 大顶堆取最大值 都是小于中位数的
// num > 中位数 小顶堆取最小值，都是大于中位数的额 如果二者长度相差大于1 则需要进行调整
type MedianFinder struct {
	maxQue maxHeap // 用于存放小于中位数的值
	minQue minHeap
}

func Constructor() MedianFinder {
	return MedianFinder{maxQue: maxHeap{}, minQue: minHeap{}}
}

func (this *MedianFinder) AddNum(num int) {
	maxHeapInMin := &this.maxQue
	minHeapInMax := &this.minQue

	if maxHeapInMin.Len() == 0 || num <= (*maxHeapInMin)[0] {
		heap.Push(maxHeapInMin, num)
		// 判断是否需要调整
		if maxHeapInMin.Len() > minHeapInMax.Len()+1 {
			value := heap.Pop(maxHeapInMin)
			heap.Push(minHeapInMax, value)
		}

	} else {
		heap.Push(minHeapInMax, num)
		// 判断是否需要调整
		if minHeapInMax.Len() > maxHeapInMin.Len()+1 {
			value := heap.Pop(minHeapInMax)
			heap.Push(maxHeapInMin, value)
		}
	}

}

func (this *MedianFinder) FindMedian() float64 {
	if this.minQue.Len() == this.maxQue.Len() {

		return (float64(this.maxQue[0]) + float64(this.minQue[0])) / 2
	}
	return float64(this.maxQue[0])
}
