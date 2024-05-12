/*
 * Copyright (c) 2024 Lei Qi. All rights reserved.
 * Author: Lei Qi
 * Description:
 * Date: 2024/2/23 下午9:34
 */

package leetcode146

type LRUCache struct {
	capacity int
	keysList []int
	keysMap  map[int]int
}

func Constructor(capacity int) LRUCache {
	//return
	return LRUCache{capacity, make([]int, 0), make(map[int]int, 0)}
}

// 如果key存在于缓存中，则返回关键字的值，否则返回-1
func (this *LRUCache) Get(key int) int {
	if ele, ok := this.keysMap[key]; ok {
		this.updateListKey(key)
		return ele
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// 关键字存在 则更新值为value
	// 不存在，则插入value
	// 如果插入超过数量capacity 则删除最久没有使用的关键字【list]
	if _, ok := this.keysMap[key]; ok {
		this.updateListKey(key)
		this.keysMap[key] = value
	} else {
		this.updateListKey(key)
		this.keysMap[key] = value
		if len(this.keysList) > this.capacity {
			delete(this.keysMap, this.keysList[0])
			this.keysList = this.keysList[1:]
		}
	}
}

func (this *LRUCache) updateListKey(key int) {
	for i := 0; i < len(this.keysList); i++ {
		if key == this.keysList[i] {
			this.keysList = append(this.keysList[:i], this.keysList[i+1:]...)
			break
		}
	}
	this.keysList = append(this.keysList, key)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
