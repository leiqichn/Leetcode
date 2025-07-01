package _0250616

import (
	"container/list"
)

// LRUCache 结构体
type LRUCache struct {
	capacity int                   // 缓存容量
	list     *list.List            // 双向链表用于维护访问顺序
	cache    map[int]*list.Element // 哈希表用于快速查找
}

// entry 用于存储键值对
type entry struct {
	key   int
	value int
}

// Constructor 初始化LRU缓存
func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		list:     list.New(),
		cache:    make(map[int]*list.Element),
	}
}

// Get 获取键对应的值
func (l *LRUCache) Get(key int) int {
	if elem, ok := l.cache[key]; ok {
		// 将访问的元素移动到链表头部表示最近使用
		l.list.MoveToFront(elem)
		return elem.Value.(*entry).value
	}
	return -1
}

// Put 插入或更新键值对
func (l *LRUCache) Put(key int, value int) {
	// 如果键已存在，更新值并移动到头部
	if elem, ok := l.cache[key]; ok {
		l.list.MoveToFront(elem)
		elem.Value.(*entry).value = value
		return
	}

	// 如果缓存已满，删除最近最少使用的元素（链表尾部）
	if l.list.Len() == l.capacity {
		tail := l.list.Back()
		if tail != nil {
			delete(l.cache, tail.Value.(*entry).key)
			l.list.Remove(tail)
		}
	}

	// 添加新元素到链表头部
	elem := l.list.PushFront(&entry{key, value})
	l.cache[key] = elem
}
