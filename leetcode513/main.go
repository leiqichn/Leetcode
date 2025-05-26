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

	res := []int{}
	queue := []*TreeNode{}

	queue = append(queue, root)
	for len(queue) != 0 {
		size := len(queue)
		levels := []int{}

		for i := 0; i < size; i++ {
			tmp := queue[0]
			queue = queue[1:]
			levels = append(levels, tmp.Val)

			if tmp.Right != nil {
				queue = append(queue, tmp.Right)
			}
			if tmp.Left != nil {
				queue = append(queue, tmp.Left)
			}

		}

		res = levels

	}
	return res[len(res)-1]
}
