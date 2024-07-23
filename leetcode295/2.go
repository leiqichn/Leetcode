/*
 * Copyright (c) 2024 Lei Qi. All rights reserved.
 * Author: Lei Qi
 * Description:
 * Date: 2024/6/3 下午11:41
 */

package leetcode295

import "container/heap"

type MinHeap []int

func (h *MinHeap) Len() int             { return len(*h) }
func (h *MinHeap) Less(i, j int) bool   { return (*h)[i] < (*h)[j] }
func (h *MinHeap) Swap(i, j int)        { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *MinHeap) Push(v interface{})   { *h = append(*h, v.(int)) }
func (h *MinHeap) Pop() (v interface{}) { *h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]; return }

type MaxHeap []int

func (h *MaxHeap) Len() int             { return len(*h) }
func (h *MaxHeap) Less(i, j int) bool   { return (*h)[i] > (*h)[j] }
func (h *MaxHeap) Swap(i, j int)        { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *MaxHeap) Push(v interface{})   { *h = append(*h, v.(int)) }
func (h *MaxHeap) Pop() (v interface{}) { *h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]; return }

/*
 *
 *                   min of bigger
 *                       ↑
 * minH (bigger  half): a_1, a_2, ..., a_m
 * maxH (smaller half): b_1, b_2, ..., b_n
 *                       ↓
 *                   max of smaller
 *
 * N = m+n is the total numbers.
 * (1). N is even, m = n = N/2
 * (2). N is odd,  m = n+1 = (N+1)/2
 *
 * AddNum(x):
 * NOTE: to make sure that all the numbers in minH >= all the numbers in maxH:
 * (1). m == n: push x to maxH, then pop the maxH, finally push it to minH.
 * (2). m != n: push x to minH, then pop the minH, finally push it to maxH.
 *
 * e.g., N = 4 -> m = n = 2; N = 5 -> m = 3, n = 2
 *
 * NOTE: actually, it doesn't matter whether minH or maxH stores ONE more number,
 * we just need always to fetch that pq which stores ONE more number.
 */

type MedianFinder struct {
	minH MinHeap // store the bigger half. NOTE: `minH.Len() - maxH.Len()` is 0 or 1.
	maxH MaxHeap // store the smaller half.
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (mf *MedianFinder) AddNum(num int) {
	minH, maxH := &mf.minH, &mf.maxH
	if minH.Len() == maxH.Len() {
		heap.Push(minH, num)
		heap.Push(maxH, heap.Pop(minH))
	} else {
		heap.Push(maxH, num)
		heap.Push(minH, heap.Pop(maxH))
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	minH, maxH := mf.minH, mf.maxH
	if minH.Len() == maxH.Len() {
		return float64(minH[0]+maxH[0]) / 2.0
	} else {
		return float64(maxH[0])
	}
}
