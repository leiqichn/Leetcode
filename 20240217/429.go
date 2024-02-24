package _0240217

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	res := [][]int{}
	queue := []*Node{}
	if root == nil {
		return [][]int{}
	}
	queue = append(queue, root)

	for len(queue) > 0 {
		size := len(queue)
		level := []int{}

		for i := 0; i < size; i++ {
			// queue top
			top := queue[0]
			queue = queue[1:]

			level = append(level, top.Val)

			for _, ele := range top.Children {
				queue = append(queue, ele)
			}
		}
		res = append(res, level)
	}
	return res
}
