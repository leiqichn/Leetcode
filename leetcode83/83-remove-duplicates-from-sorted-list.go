/*
 * Copyright (c) 2024 Lei Qi. All rights reserved.
 * Author: Lei Qi
 * Description:
 * Date: 2024/1/14 下午5:51
 */

package leetcode83

import . "lcutils/listnode"

func deleteDuplicates(head *ListNode) *ListNode {
	for i := head; i != nil && i.Next != nil; i = i.Next {
		if i.Next.Val == i.Val {
			i.Next = i.Next.Next
		} else {
			i = i.Next
		}
	}
	return head
}
