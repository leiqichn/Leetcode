package _0251012

func findCircleNum(isConnected [][]int) int {

	vis := make([]bool, len(isConnected))

	var dfs func(int)

	dfs = func(from int) {
		if vis[from] {
			return
		}

		vis[from] = true

		for to, conn := range isConnected[from] {
			if conn == 1 {
				dfs(to)
			}
		}

	}
	ans := 0
	for i, v := range vis {
		if !v {
			ans++
			dfs(i)
		}
	}

	return ans

}
