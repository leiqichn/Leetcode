package _0250802_map_intlist

type LRUCache struct {
	// intlist capacity n
	capacity  int
	orderList []*node
	cacheMaps map[int]*node
}

type node struct {
	key int
	val int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity:  capacity,
		orderList: make([]*node, 0),
		cacheMaps: make(map[int]*node),
	}
}

func (this *LRUCache) Put(key int, value int) {
	// is eist ?
	if elem, ok := this.cacheMaps[key]; ok {
		// 更新值
		elem.val = value
		// 挪到最后
		this.Update(elem.key)
		return // 提前返回
	}
	if len(this.orderList) >= this.capacity {
		// 两次delete
		oldest := this.orderList[0]
		this.orderList = this.orderList[1:]
		delete(this.cacheMaps, oldest.key)
	}

	// add
	newNode := &node{
		key: key,
		val: value,
	}
	this.orderList = append(this.orderList, newNode)
	this.cacheMaps[key] = newNode
}

func (this *LRUCache) Update(key int) {
	// 挪到最后list
	for i, elem := range this.orderList {
		if key == elem.key {
			this.orderList = append(this.orderList[:i], append(this.orderList[:i+1], elem)...) // this.orderList[:i+1] 这里错误， 应为this.orderList[i+1:]
			return
		}
	}
}

func (this *LRUCache) Get(key int) int {

	if elem, ok := this.cacheMaps[key]; ok {
		this.Update(elem.key)
		return elem.val
	}
	return -1
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
