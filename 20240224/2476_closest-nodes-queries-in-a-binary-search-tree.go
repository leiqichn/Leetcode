/*
 * Copyright (c) 2024 Lei Qi. All rights reserved.
 * Author: Lei Qi
 * Description:
 * Date: 2024/2/24 下午10:11
 */

package _0240224 /**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestNodes(root *TreeNode, queries []int) [][]int {
	// 遍历二叉树  将结果保存至list 中 找到最大最小值
	ans := [][]int{}
	res := []int{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		res = append(res, root.Val)
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	sort.Ints(res)
	for _, item := range queries {
		tmpMin := []int{}
		tmpMax := []int{}
		for _, eleTree := range res {
			if eleTree < item {
				tmpMin = append(tmpMin, eleTree)
			} else if eleTree > item {
				tmpMax = append(tmpMax, eleTree)
			} else {
				tmpMin = append(tmpMin, eleTree)
				tmpMax = append(tmpMax, eleTree)
			}
		}
		tmpRes := []int{}

		if len(tmpMin) == 0 {
			tmpRes = append(tmpRes, -1)
		} else {
			tmpRes = append(tmpRes, tmpMin[len(tmpMin)-1])
		}

		if len(tmpMax) == 0 {
			tmpRes = append(tmpRes, -1)
		} else {
			tmpRes = append(tmpRes, tmpMax[0])
		}
		ans = append(ans, tmpRes)

	}
	return ans
}
