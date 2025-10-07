package _0251007

type LRUCache struct {
	capacity int
	maps     map[int]int
	lists    []*node
}

type node struct {
	key int
	val int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		maps:     make(map[int]int),
		lists:    make([]*node, 0),
	}
}

func (this *LRUCache) Put(key int, value int) {
	// 如果存在， 直接更新value, 并返回

	if _, ok := this.maps[key]; ok {
		this.maps[key] = value
		this.Update(key)
		return
	}

	// 判断长度是否超过capacity ， 同步删除两个最list【0】

	if len(this.lists) >= this.capacity {
		old := this.lists[0]
		this.lists = this.lists[1:]
		delete(this.maps, old.key)
	}

	// 不存在， 则插入

	newNode := &node{
		key: key,
		val: value,
	}
	this.lists = append(this.lists, newNode)
	this.maps[key] = value

}

// 更新位置到最新上 list 最新的位置
func (this *LRUCache) Update(key int) {
	for i := 0; i < len(this.lists); i++ {
		if key == this.lists[i].key {
			tmp := this.lists[i]
			this.lists = append(this.lists[:i], append(this.lists[i+1:], tmp)...)
			return
		}
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.maps[key]; ok {

		this.Update(key)
		return v
	}
	return -1
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
