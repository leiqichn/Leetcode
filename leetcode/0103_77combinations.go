package leetcode

var (
	path []int
	res  [][]int
)

func combine(n int, k int) [][]int {
	path, res = make([]int, 0, k), make([][]int, 0)
	dfs(n, k, 1)
	return res
}

func dfs(n int, k int, startIndex int) {
	if len(path) == k {
		// 终止条件
		var tmp []int
		tmp = make([]int, k)
		copy(tmp, path)
		res = append(res, tmp)
		return
	}
	// 每一个循环里的操作
	for i := startIndex; i <= n; i++ {
		//操作 将结果添加到path
		path = append(path, i)
		dfs(n, k, i+1)
		//
		path = path[:len(path)-1]
	}

}
