package _0251013

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

func maxDepth(root *TreeNode) int {
	ans := 0
	var dfs func(root *TreeNode, dep int)
	dfs = func(root *TreeNode, dep int) {
		if root == nil {
			return
		}
		dep++
		ans = max(ans, dep)
		dfs(root.Right, dep)
		dfs(root.Left, dep)

	}

	dfs(root, 0)
	return ans
}
