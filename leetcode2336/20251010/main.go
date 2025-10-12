package _0251010

import "container/heap"

type SmallestInfiniteSet struct {
	// 最小堆
	heap *IntHeap
	//map
	maps map[int]bool
}

func Constructor() SmallestInfiniteSet {
	set := SmallestInfiniteSet{
		heap: &IntHeap{},
		maps: make(map[int]bool),
	}

	heap.Init(set.heap)

	return set
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	return heap.Pop(this.heap).(int)
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if _, ok := this.maps[num]; ok {
		return
	}
	heap.Push(this.heap, num)
	this.maps[num] = true
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
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

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */
