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

// 递归 是有效的搜索树 后序遍历 因为这个左边的最大值的来源可能是左边子树的右子树 所以要反馈一个区间，而不是左边返回最大值，右边返回最小值

func dfs(node *TreeNode) (int, int) {
	if node == nil {
		return math.MaxInt, math.MinInt
	}
	lMin, lMax := dfs(node.Left)
	rMin, rMax := dfs(node.Right)
	x := node.Val
	// 也可以在递归完左子树之后立刻判断，如果发现不是二叉搜索树，就不用递归右子树了
	if x <= lMax || x >= rMin {
		return math.MinInt, math.MaxInt // 不是的条件
	}
	return min(lMin, x), max(rMax, x)
}

func isValidBST(root *TreeNode) bool {
	_, mx := dfs(root)
	return mx != math.MaxInt
}
