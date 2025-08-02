package _0250802

import (
	"container/list"
)

type LRUCache struct {
	// 双向列表保存顺序，map进行快速查询
	capacity int
	list     list.List
	maps     map[int]*list.Element
}

type node struct {
	key int
	val int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		list:     list.List{},
		maps:     make(map[int]*list.Element),
	}
}

func (this *LRUCache) Put(key int, value int) {
	// if in move to back

	if _, ok := this.maps[key]; ok {
		ele := this.maps[key]
		ele.Value.(*node).val = value
		this.list.MoveToBack(ele)
		return
	} else {
		// not in 1. cap>  delete(maps list) add
		if this.list.Len() >= this.capacity {
			lastElem := this.list.Front()
			delete(this.maps, lastElem.Value.(*node).key) // delete maps don't forget
			this.list.Remove(lastElem)
		}
	}

	// 2. add
	newNode := this.list.PushBack(&node{
		key: key,
		val: value,
	}) // 这里存储的指针， 所以 ele.Value.(*node).val 需要使用*node
	this.maps[key] = newNode
}

func (this *LRUCache) Get(key int) int {
	// 查询， 如果存在 则返回关键字， 并且挪到最后
	if _, ok := this.maps[key]; ok {
		ele := this.maps[key]
		this.list.MoveToBack(ele)
		return ele.Value.(*node).val
	}
	return -1
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
