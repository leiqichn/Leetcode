package MyselfListNode

// 手动实现双向列表
type LRUCache struct {
	capacity int
	head     *Node // 虚拟头节点
	tail     *Node // 虚拟尾节点
	cache    map[int]*Node
}

// 链表节点
type Node struct {
	key  int
	val  int
	prev *Node
	next *Node
}

func Constructor(capacity int) LRUCache {
	// 初始化虚拟头尾节点（简化边界判断）
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return LRUCache{
		capacity: capacity,
		head:     head,
		tail:     tail,
		cache:    make(map[int]*Node),
	}
}

// 从链表中移除节点
func (this *LRUCache) _removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// 添加节点到链表尾部（表示最近使用）
func (this *LRUCache) _addToTail(node *Node) {
	// node 前后
	node.prev = this.tail.prev
	node.next = this.tail
	// node 前一个节点前后
	this.tail.prev.next = node
	this.tail.prev = node
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this._removeNode(node) // 从当前位置移除
		this._addToTail(node)  // 移到尾部
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// 如果 key 已存在，更新值并移到尾部
	if node, ok := this.cache[key]; ok {
		node.val = value
		this._removeNode(node)
		this._addToTail(node)
		return
	}

	// 如果缓存已满，删除头部的下一个节点（最久未使用）
	if len(this.cache) >= this.capacity {
		oldest := this.head.next
		this._removeNode(oldest)
		delete(this.cache, oldest.key)
	}

	// 创建新节点并添加到尾部
	newNode := &Node{key: key, val: value}
	this.cache[key] = newNode
	this._addToTail(newNode)
}
