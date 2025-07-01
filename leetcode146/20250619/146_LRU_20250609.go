package leetcode146

import "container/list"

type LRUCache struct {
	hash      map[int]*list.Element
	orderList *list.List
	cap       int
}

type entry struct {
	key int
	val int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		hash:      make(map[int]*list.Element, capacity),
		orderList: list.New(),
		cap:       capacity,
	}

}

func (this *LRUCache) Get(key int) int {
	if elem, ok := this.hash[key]; ok {
		this.orderList.MoveToBack(elem)
		return this.hash[key].Value.(*entry).val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// 1 key 存在
	if elem, ok := this.hash[key]; ok {
		elem.Value.(*entry).val = value
		this.orderList.MoveToBack(this.hash[key])
		return // 不要漏了return
	}

	// 2 需要老化
	if len(this.hash) == this.cap {
		front := this.orderList.Front()
		delete(this.hash, front.Value.(*entry).key)
		this.orderList.Remove(front)
	}

	// 3 新增Key
	newElem := this.orderList.PushBack(&entry{
		key: key,
		val: value,
	})
	this.hash[key] = newElem
}
