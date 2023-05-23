package leetcode

import "container/list"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	stack := list.New()
	res := []int{}

	if root == nil { //防止为空
		return res
	}
	stack.PushBack(root)
	for stack.Len() > 0 {
		top := stack.Remove(stack.Back()).(*TreeNode)
		res = append(res, top.Val)
		if top.Right != nil {
			stack.PushBack(top.Right)
		}
		if top.Left != nil {
			stack.PushBack(top.Left)
		}
	}
	return res
}
