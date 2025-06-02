package leetcode235

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	//if root == nil { 这个部分不用考虑，因为题目中说明了 p,q 一定存在
	//	return nil
	//}

	// 如果root val 都大于两个值，说明都在左子树
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}

	// 如果root val 都小于于两个值右边
	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	return root
}
