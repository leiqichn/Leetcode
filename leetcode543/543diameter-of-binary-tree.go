/*
 * Copyright (c) 2024 Lei Qi. All rights reserved.
 * Author: Lei Qi
 * Description:
 * Date: 2024/5/8 上午12:55
 */

package leetcode543

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) (ans int) {
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1 // 下面 +1 后，对于叶子节点就刚好是 0
		}
		lLen := dfs(node.Left) + 1  // 左子树最大链长+1
		rLen := dfs(node.Right) + 1 // 右子树最大链长+1
		ans = max(ans, lLen+rLen)   // 两条链拼成路径
		return max(lLen, rLen)      // 当前子树最大链长
	}
	dfs(root)
	return
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
