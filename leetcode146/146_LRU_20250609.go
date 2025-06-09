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
		return this.hash[key].Value.(entry).val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// 1 key 存在
	if _, ok := this.hash[key]; ok {
		// this.hash[key].Value.(*entry).val = value // 错误！！！！ 值不能直接访问， A[值类型存储] --> B[修改时需整体替换]  C[指针类型存储] --> D[可直接修改字段]
		this.hash[key].Value = entry{
			key: key,
			val: value,
		}
		this.orderList.MoveToBack(this.hash[key])
	}

	// 2 需要老化
	if len(this.hash) == this.cap {
		front := this.orderList.Front()
		delete(this.hash, front.Value.(entry).key)
		this.orderList.Remove(front)
	}

	// 3 新增Key
	newElem := this.orderList.PushBack(entry{
		key: key,
		val: value,
	})
	this.hash[key] = newElem
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
