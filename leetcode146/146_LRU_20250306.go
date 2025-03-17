package leetcode146

type LRUCache struct {
	orderList []int
	cacheMaps map[int]int
	capacity  int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		orderList: make([]int, 0, capacity), // 39 分钟调试完成，需要确认的是，List初始化需要为空，而不是orderList: make([]int, capacity) 这个是创建capacity个0
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
	} else if len(this.cacheMaps) >= this.capacity {
		// 删除key
		oldKey := this.orderList[0]
		this.orderList = this.orderList[1:]
		delete(this.cacheMaps, oldKey)
		// 新建key
		this.cacheMaps[key] = value
		this.orderList = append(this.orderList, key)
	} else if len(this.cacheMaps) < this.capacity {
		// 新建key
		this.cacheMaps[key] = value
		this.orderList = append(this.orderList, key)
	}
}
