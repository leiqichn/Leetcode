/*
 * Copyright (c) 2024 Lei Qi. All rights reserved.
 * Author: Lei Qi
 * Description:
 * Date: 2024/5/10 下午7:41
 */

package leetcode230

import . "lcutils/treenode"

// 转化为list 取index == k-1 的元素
// 中序遍历还原顺序
func kthSmallest(root *TreeNode, k int) int {
	res := []int{}
	var dfs func(root *TreeNode) *TreeNode
	dfs = func(root *TreeNode) *TreeNode {
		if root == nil {
			return root
		}

		dfs(root.Left)
		res = append(res, root.Val)
		dfs(root.Left)
		return root
	}
	dfs(root)

	return res[k-1]

}
