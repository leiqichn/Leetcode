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

	return helper2(root, &pre)
}

// 中序遍历
func helper2(root *TreeNode, pre *int) bool {

	if root == nil {
		return true
	}

	if !helper2(root.Left, pre) {
		return false
	}

	if root.Val < *pre {
		return false
	}

	*pre = root.Val // 只改变指针的值，而不是地址 pre = &root.Val

	return helper2(root.Right, pre)
}

func isValidBST3(root *TreeNode) bool {
	pre := math.MinInt64 // 使用闭包共享变量pre
	var dfs func(*TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if !dfs(node.Left) {
			return false
		} // 左
		if node.Val <= pre {
			return false
		} // 中
		pre = node.Val
		if !dfs(node.Right) { //右
			return false
		}
		return true
	}
	return dfs(root)
}
