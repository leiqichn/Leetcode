package leetcode

import "fmt"

type MyHashSet struct {
	hashSet map[int]int
}

func Constructor() MyHashSet {
	return MyHashSet{hashSet: make(map[int]int)}
}

func (this *MyHashSet) Add(key int) {
	if _, ok := this.hashSet[key]; ok {
		this.hashSet[key]++
	} else {
		this.hashSet[key] = 1
	}
}

func (this *MyHashSet) Remove(key int) {
	if _, ok := this.hashSet[key]; !ok {
		fmt.Println("err")
	}
	delete(this.hashSet, key)
}

func (this *MyHashSet) Contains(key int) bool {
	if _, ok := this.hashSet[key]; ok {
		return true
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
