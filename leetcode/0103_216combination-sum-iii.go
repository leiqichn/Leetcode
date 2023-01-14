package leetcode

import "fmt"

// 剪枝
//var (
//	path []int
//	res  [][]int
//)

func combinationSum3(n int, k int) [][]int {
	path, res = make([]int, 0, k), make([][]int, 0)
	dfs2(n, k, 1)
	return res
}

func sum(path []int) int {
	var sumRes int
	for _, val := range path {
		sumRes += val
	}
	return sumRes
}

func dfs2(n int, k int, startIndex int) {
	if len(path) == k && sum(path) == n {
		// 终止条件
		var aname int
		aname = sum(path)
		fmt.Print(aname)
		var tmp []int
		tmp = make([]int, k)
		copy(tmp, path) // copy 函数必须要保证前面是目标 ，后边是源，两个长度是一样的，这样才能复制成功，所以tmp 必须make
		res = append(res, tmp)
		return
	}
	// 每一个循环里的操作
	// 减枝；for 循环的边界，n-i（剩下的元素）>= k- len(path) (还缺少的元素)，算出 i 至少从多少开始
	for i := startIndex; i <= 9-(k-len(path))+1; i++ { // 注意i 从start 开始 // 减枝优化
		//操作 将结果添加到path
		path = append(path, i)
		dfs2(n, k, i+1)
		// 撤回操作
		path = path[:len(path)-1]
	}

}

var (
	path []int
	res  [][]int
)

func combinationSum3(n int, k int) [][]int {
	path, res = make([]int, 0, k), make([][]int, 0)
	dfs3(n, k, 1, 0)
	return res
}

func dfs3(n int, k int, startIndex int, sum int) {
	if len(path) == k {
		if sum == n {
			var tmp []int
			tmp = make([]int, k)
			copy(tmp, path)
			res = append(res, tmp)
		}
		return
	}
	for i := startIndex; i <= 9; i++ { // 注意i 从start 开始 减枝优化
		//操作将结果添加到path
		sum += i
		path = append(path, i)
		dfs3(n, k, i+1, sum)
		// 撤回操作
		sum -= i
		path = path[:len(path)-1]

	}

}
