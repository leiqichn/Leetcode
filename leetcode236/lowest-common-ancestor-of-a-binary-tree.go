package leetcode236

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import . "lcutils/treenode"

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 当前节点为空，当前节点为p, 当前节点为q
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	// 其他：左右子树都找到， 返回当前节点
	if left != nil && right != nil {
		return root
	}
	// 只有左子树找到
	if left != nil {
		return left
	}
	// 只有右子树找到
	if right != nil {
		return right
	}
	// 左右都没找到
	return nil
}
