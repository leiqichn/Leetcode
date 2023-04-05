package leetcode

func sumOfLeftLeaves(root *TreeNode) int {
	// 递归的输入和返回值是对的了
	res := 0
	// 终止条件
	if root == nil {
		return 0
	}
	leftNode := root.Left
	if leftNode != nil && leftNode.Left == nil && leftNode.Right == nil {
		res += leftNode.Val
	}
	// 单次循环
	return res
}
