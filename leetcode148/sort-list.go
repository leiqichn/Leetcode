/*
 * Copyright (c) 2024 Lei Qi. All rights reserved.
 * Author: Lei Qi
 * Description:
 * Date: 2024/7/23 下午11:36
 */

package leetcode148

import "sort"
import . "lcutils/listnode"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	// 使用list 排序，然后重排链表
	if head == nil {
		return nil
	}

	listNodes := []*ListNode{}

	for i := head; i != nil; i = i.Next {
		listNodes = append(listNodes, i)
	}
	sort.Slice(listNodes, func(i, j int) bool {
		return listNodes[i].Val < listNodes[j].Val
	})

	for i := 1; i < len(listNodes); i++ {
		listNodes[i-1].Next = listNodes[i]
	}
	listNodes[len(listNodes)-1].Next = nil // 需要设置最后的nil
	return listNodes[0]
}
