package leetcode

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	preorder(root, res)
	return res
}

func preorder(root *TreeNode, res []int) {
	if root == nil {
		return
	}

	res = append(res, root.Val)
	preorder(root.Left, res)
	preorder(root.Right, res)

}

// 中序遍历 递归
func inorderTraversal(root *TreeNode) []int {
	res := []int{}

	st := list.New()
	cur := root
	if cur == nil {
		return res
	}

	for cur != nil || st.Len() > 0 {
		if cur != nil {
			st.PushBack(cur)
			cur = cur.Left
		} else {
			top := st.Remove(st.Back()).(*TreeNode)
			res = append(res, top.Val)
			cur = cur.Right
		}
	}
	return res
}

// 前序遍历 栈法
func preorderTraversal2(root *TreeNode) []int {
	res := []int{}

	st := list.New()
	cur := root
	if cur == nil {
		return res
	}
	st.PushBack(root)
	for cur != nil {
		cur := st.Remove(st.Back()).(*TreeNode)
		res = append(res, cur.Val)
		if cur.Right != nil {
			st.PushBack(cur.Right)
		}

		if cur.Left != nil { // 空节点不入栈，
			st.PushBack(cur.Left)
		}
	}
	return res
}
