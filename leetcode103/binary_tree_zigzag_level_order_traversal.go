package leetcode103

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import (
	. "lcutils/treenode"
	"sort"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := [][]int{}
	queue := []*TreeNode{}

	queue = append(queue, root)
	even := false

	for len(queue) != 0 {
		size := len(queue)
		levels := []int{}

		for i := 0; i < size; i++ {
			tmp := queue[0]
			queue = queue[1:]
			levels = append(levels, tmp.Val)

			if tmp.Left != nil {
				queue = append(queue, tmp.Left)
			}

			if tmp.Right != nil {
				queue = append(queue, tmp.Right)
			}

		}
		// 偶数需要反转
		if even {
			sort.Sort(sort.Reverse(sort.IntSlice(levels)))
		}
		res = append(res, levels)
		// 每次都变化
		even = !even
	}
	return res
}

func revert(list []int) {

	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}

}
