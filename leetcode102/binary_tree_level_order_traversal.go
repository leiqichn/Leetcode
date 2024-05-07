/*
 * Copyright (c) 2024 Lei Qi. All rights reserved.
 * Author: Lei Qi
 * Description:
 * Date: 2024/5/8 上午12:55
 */

package leetcode102


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

* type TreeNode struct {
	    Val int
	    Left *TreeNode
	   Right *TreeNode
	 }

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	queue := []*TreeNode {}

	if root == nil {
		return res
	}

	queue = append(queue, root)

	for len(queue) >0 {
		size := len(queue)
		level := []int{}
		for i:=0;i <size; i++ {
			top := queue[0]
			queue = queue[1:len(queue)]
			level = append(level, top.Val)

			if top.Right != nil {
				queue = append(queue, top.Right)
			}
			if top.Left != nil {
				queue = append(queue ,top.Left)
			}
		}
		res = append(res,level)
	}
	return res
}