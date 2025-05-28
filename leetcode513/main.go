package leetcode513

import . "lcutils/treenode"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
	queue := []*TreeNode{}
	node := root
	queue = append(queue, root)
	for len(queue) != 0 {
		tmp := queue[0]
		queue = queue[1:]
		node = tmp
		if tmp.Right != nil {
			queue = append(queue, tmp.Right)
		}
		if tmp.Left != nil {
			queue = append(queue, tmp.Left)
		}
	}
	return node.Val
}
