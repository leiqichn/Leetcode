package leetcode

// 迭代法 使用队列
func maxDepth2(root *TreeNode) int {
	depth := 0
	if root == nil {
		return 0
	}
	queue := []*TreeNode{}

	queue = append(queue, root)
	for len(queue) > 0 { // queue 不为空的时候
		size := len(queue)
		for i := 0; i < size; i++ { // 遍历一层
			top := queue[0]
			queue = queue[1:] // 取最上层元素，并切掉该元素
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
		}
		depth++
	}
	return depth
}
