/*
 * Copyright (c) 2024 Lei Qi. All rights reserved.
 * Author: Lei Qi
 * Description:
 * Date: 2024/5/8 下午11:32
 */

package leetcode108

import . "lcutils/treenode"

func sortedArrayToBST(nums []int) *TreeNode {
	var travesal func(list []int, left, right int) *TreeNode
	// 递归的定义：返回root 节点，
	travesal = func(list []int, left, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := (left + right) / 2
		root := &TreeNode{Val: list[mid]}
		root.Left = travesal(nums, left, mid-1)
		root.Right = travesal(nums, mid+1, right)
		return root
	}

	return travesal(nums, 0, len(nums)-1)
}
