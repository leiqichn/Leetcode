/*
 * Copyright (c) 2024 Lei Qi. All rights reserved.
 * Author: Lei Qi
 * Description:
 * Date: 2024/5/9 下午11:17
 */

package leetcode98

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import (
	. "lcutils/treenode"
	"math"
)

// 递归 是有效的搜索树
// 中序遍历, 通过和前一个值比较如果小于前一个值，就返回
func isValidBST2(root *TreeNode) bool {
	pre := math.MinInt64

	return helper(root, &pre)
}

// 中序遍历
func helper(root *TreeNode, pre *int) bool {

	if root == nil {
		return true
	}

	if !helper(root.Left, pre) {
		return false
	}

	if root.Val < *pre {
		return false
	}

	*pre = root.Val // 只改变指针的值，而不是地址 pre = &root.Val

	return helper(root.Right, pre)
}
