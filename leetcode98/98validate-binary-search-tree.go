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
// 根节点传入的时候，是没有范围的，需要传入正无穷和负无穷
func isValidBST4(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

// 前序遍历
func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	// 提前返回不合规的条件
	if root.Val <= lower || root.Val >= upper {
		return false
	} // else 情况下也不能直接返回true, 而是要递归的看其子树
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}
