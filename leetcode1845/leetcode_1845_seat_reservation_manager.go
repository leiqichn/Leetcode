/*
 * Copyright (c) 2023 Lei Qi. All rights reserved.
 * Author: Lei Qi
 * Description: leetcode 1825 设计题
 * Date: 2023/10/5 下午4:39
 */

package leetcode1845

import (
	"fmt"
	"sort"
)

type SeatManager struct {
	seatMap   map[int]int
	seatSlice []int
	minSeat   int
}

func Constructor(n int) SeatManager {
	seatMap := make(map[int]int, n+1)
	seatSlice := []int{} // 和make([]int, n+1) 区别
	for i := 1; i <= n; i++ {
		seatMap[i] = 0
		seatSlice = append(seatSlice, i)
	}
	seatMan := SeatManager{
		seatMap,
		seatSlice,
		1,
	}
	return seatMan
}

func (this *SeatManager) Reserve() int {
	//　排序
	sort.Slice(this.seatSlice, func(i, j int) bool {
		return this.seatSlice[i] < this.seatSlice[j] // 排序的规则
	})
	fmt.Println(this.seatSlice)
	this.minSeat = this.seatSlice[0]
	this.seatSlice = this.seatSlice[1:]
	this.seatMap[this.minSeat] = 1
	return this.minSeat
}

func (this *SeatManager) Unreserve(seatNumber int) {
	if this.seatMap[seatNumber] == 1 { // 在复制前边的代码的时候一点要看清楚了
		this.seatMap[seatNumber] = 0
		this.seatSlice = append(this.seatSlice, seatNumber)
	}
}

/**
 * Your SeatManager object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Reserve();
 * obj.Unreserve(seatNumber);
 */
