package leetcode

// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码还未经过力扣测试，仅供参考，如有疑惑，可以参照我写的 java 代码对比查看。

import "container/list"

type LRUCache struct {
	cap   int                   // 缓存容量
	cache map[int]*list.Element // 双向链表节点 指向的map
	list  *list.List            // 双向链表
}

type keyVal struct {
	key, val int // 节点的Key和Value
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:   capacity,                    // 初始化缓存容量
		cache: make(map[int]*list.Element), // 初始化map映射
		list:  list.New(),                  // 初始化双向链表
	}
}

func (this *LRUCache) Get(key int) int {
	if elem, ok := this.cache[key]; ok { // 如果map里有key对应的双向链表节点
		this.list.MoveToFront(elem)     // 把节点移动到链表头
		return elem.Value.(*keyVal).val // 返回节点的value值
	}
	return -1 // 没有找到的情况下，返回-1
}

func (this *LRUCache) Put(key int, value int) {
	if elem, ok := this.cache[key]; ok { // 如果map里有key对应的双向链表节点
		this.list.MoveToFront(elem)      // 把节点移动到链表头
		elem.Value.(*keyVal).val = value // 更新节点的value值
		return
	}
	if this.list.Len() >= this.cap { // 如果超过了缓存容量
		tail := this.list.Back()      // 获取链表的尾节点
		k := tail.Value.(*keyVal).key // 获取节点的key
		this.list.Remove(tail)        // 从链表中删除尾节点
		delete(this.cache, k)         // 从map中删除尾节点
	}
	elem := this.list.PushFront(&keyVal{key, value}) // 将节点添加到链表头
	this.cache[key] = elem                           // 将节点映射到map中
}
