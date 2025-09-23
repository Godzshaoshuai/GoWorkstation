package FastCode

// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。dfs
func numIslands(grid [][]byte) (ans int) {
	m, n := len(grid), len(grid[0])

	var dfs func(int, int)
	dfs = func(i int, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != '1' {
			return
		}
		grid[i][j] = '2'
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}
	for i, row := range grid {
		for j, col := range row {
			if col == '1' {
				dfs(i, j)
				ans++
			}
		}
	}
	return
}
