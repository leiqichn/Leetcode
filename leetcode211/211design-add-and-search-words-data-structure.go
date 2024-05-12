/*
 * Copyright (c) 2024 Lei Qi. All rights reserved.
 * Author: Lei Qi
 * Description:
 * Date: 2024/5/8 上午12:55
 */

package leetcode211

import "strings"

type WordDictionary struct {
	wordMap map[string]bool
}

func Constructor() WordDictionary {
	return WordDictionary{make(map[string]bool, 0)}
}

func (this *WordDictionary) AddWord(word string) {
	this.wordMap[word] = true
}

func (this *WordDictionary) Search(word string) bool {
	if _, ok := this.wordMap[word]; ok {

		return true
	} else if strings.Contains(word, ".") {
		for key, _ := range this.wordMap {
			if len(key) == len(word) {
				flag := true
				for i := 0; i < len(word); i++ {
					if key[i] == word[i] || word[i] == '.' {

					} else {
						flag = false
					}
				}
				if flag == true {
					return true
				}
			}
		}
		return false
	}
	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
