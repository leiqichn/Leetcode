package leetcode146_20250306

type LRUCache struct {
	orderList []int
	cacheMaps map[int]int
	capacity  int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		orderList: make([]int, capacity),
		cacheMaps: make(map[int]int, capacity),
		capacity:  capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	// key 存在与单独的环境中
	if _, ok := this.cacheMaps[key]; ok {
		this.Update(key)
		return this.cacheMaps[key]
	}
	return -1
}

func (this *LRUCache) Update(key int) {
	// 更新key 到最新位置
	for i := 0; i < len(this.orderList); i++ {
		if this.orderList[i] == key {
			this.orderList = append(this.orderList[:i], append(this.orderList[i+1:], this.orderList[i])...)
		}
	}
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cacheMaps[key]; ok {
		// 更新key
		this.cacheMaps[key] = value
		this.Update(key)
	} else if len(this.cacheMaps) <= this.capacity-1 {
		// 新建key
		this.cacheMaps[key] = value
		this.orderList = append(this.orderList, key)
	} else if len(this.cacheMaps) >= this.capacity {
		// 删除key
		oldKey := this.orderList[0]
		this.orderList = this.orderList[1:]
		delete(this.cacheMaps, oldKey)
		// 新建key
		this.cacheMaps[key] = value
		this.orderList = append(this.orderList, key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//leetcode submit region end(Prohibit modification and deletion)
