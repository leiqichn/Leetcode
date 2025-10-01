package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
	res := -1
	queue := []*TreeNode{}

	if root == nil {
		return -1
	}

	queue = append(queue, root)
	//res 在哪里更新呢？
	for len(queue) > 0 {
		size := len(queue)
		// 遍历每一层
		for i := 0; i < size; i++ {
			top := queue[0]
			queue = queue[1:]
			// 获取每一层的最左边的位置，更新res
			if i == 0 {
				res = top.Val
			}
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
		}

	}
	return res
}

// 不需要最大值，只需要最后一行最最左边的元素
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
