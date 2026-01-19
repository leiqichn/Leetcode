package _0260119

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

type node struct {
	level int
	sum   int
}

func maxLevelSum(root *TreeNode) int {
	res := node{100000, -10000000}
	if root == nil {
		return 0
	}

	queue := []*TreeNode{}

	queue = append(queue, root)
	level := 1
	for len(queue) > 0 {
		size := len(queue)
		sum := 0

		for i := 0; i < size; i++ {
			top := queue[0]
			queue = queue[1:]
			sum += top.Val
			if top.Left != nil {
				queue = append(queue, top.Left)
			}

			if top.Right != nil {
				queue = append(queue, top.Right)
			}

		}
		if sum > res.sum || (sum == res.sum && level < res.level) {
			res = node{level, sum}
		}

		level++
	}

	return res.level

}
