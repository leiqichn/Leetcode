package leetcode104

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

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{}

	queue = append(queue, root)
	deeps := 0
	for len(queue) > 0 {
		size := len(queue)
		deeps++
		// nextLevel := []*TreeNode{}

		for i := 0; i < size; i++ {
			top := queue[0]
			queue = queue[1:]

			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
		}
	}
	return deeps
}
