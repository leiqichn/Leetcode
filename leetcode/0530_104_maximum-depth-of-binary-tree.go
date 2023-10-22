package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "container/list"

func maxDepth(root *TreeNode) int {
	depth := 0
	if root == nil {
		return 0
	}
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			// 切掉第一个元素
			first := queue.Remove(queue.Front()).(*TreeNode)
			if first.Left != nil {
				queue.PushBack(first.Left)
			}
			if first.Right != nil {
				queue.PushBack(first.Right)
			}
		}
		depth++
	}
	return depth
}
