package leetcode114

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
)

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	if root.Right == nil && root.Left == nil {
		return
	}

	res := front(root)
	for _, treenode := range res {
		if treenode != nil {
			root.Right = treenode
			root.Left = nil
			root = root.Right // 指定root 变换
		} else {
			continue
		}
	}
}

func front(root *TreeNode) (res []*TreeNode) {
	if root == nil {
		return
	}
	res = append(res, root) //res 需要全局变量
	front(root.Left)
	front(root.Right)
	return res
}

// 正确解法

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	if root.Right == nil && root.Left == nil {
		return
	}

	res := []*TreeNode{} // res 在外边

	var help func(root *TreeNode)
	help = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root)
		help(root.Left)
		help(root.Right)
		return
	}
	help(root)
	for _, treenode := range res {
		if treenode != nil {
			root.Right = treenode
			root.Left = nil
			root = root.Right
		} else {
			continue
		}
	}
}
