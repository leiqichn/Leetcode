package leetcode

import "container/list"

// 链地址法
var base = 769

// base := 769
type entry struct {
	k int
	v int
}

type MyHashMap struct {
	hashMaps []list.List
}

// hash集合可以使用数组链表；
func Constructor() MyHashMap {
	return MyHashMap{make([]list.List, base)}
}

// func　(this *MyHashMap) hash(key int, value int){

// }

func (this *MyHashMap) Put(key int, value int) {
	hash := key % base
	for e := this.hashMaps[hash].Front(); e != nil; e = e.Next() {
		if et := e.Value.(entry); et.k == key {
			et = entry{key, value}
			return
		}
	}
	this.hashMaps[hash].PushBack(entry{key, value})
}

func (this *MyHashMap) Get(key int) int {
	hash := key % base
	for e := this.hashMaps[hash].Front(); e != nil; e = e.Next() {
		if et, ok := e.Value.(entry); ok {
			return et.v
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	hash := key % base
	for e := this.hashMaps[hash].Front(); e != nil; e = e.Next() {
		if et, ok := e.Value.(entry); ok && et.k == key {
			this.hashMaps[hash].Remove(e)
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */

// 链地址法
var base = 769

// base := 769
type entry struct {
	k int
	v int
}

type MyHashMap struct {
	hashMaps []list.List
}

// hash集合可以使用数组链表；
func Constructor() MyHashMap {
	return MyHashMap{make([]list.List, base)}
}

// func　(this *MyHashMap) hash(key int, value int){

// }

func (this *MyHashMap) Put(key int, value int) {
	hash := key % base
	for e := this.hashMaps[hash].Front(); e != nil; e = e.Next() {
		if et := e.Value.(entry); et.k == key {
			et.v = value // 注意这里et 是拷贝，需要是使用e.Value 而不是其拷贝
			return
		}
	}
	this.hashMaps[hash].PushBack(&entry{key, value})
}

func (this *MyHashMap) Get(key int) int {
	hash := key % base
	for e := this.hashMaps[hash].Front(); e != nil; e = e.Next() {
		if et, ok := e.Value.(entry); ok && et.k == key { // 可以缩写为if et := e.Value.(entry); et.k == key  用于断言，并找key
			return et.v
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	hash := key % base
	for e := this.hashMaps[hash].Front(); e != nil; e = e.Next() {
		if et, ok := e.Value.(entry); ok && et.k == key {
			this.hashMaps[hash].Remove(e)
			return
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
