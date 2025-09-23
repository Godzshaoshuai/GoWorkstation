package main

import "fmt"

func main() {
	var n, m, start1, final1, a, b, d, e, f int
	fmt.Scanln(&n, &m)
	fmt.Scanln(&start1, &final1)
	fmt.Scanln(&a, &b)
	nums := make([][]int, n)
	visit := make([][]bool, n)
	for i := 0; i < n; i++ {
		nums[i] = make([]int, n)
		visit[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		fmt.Scanln(&d, &e, &f)
		nums[d-1][e-1] = f
		nums[e-1][d-1] = f
	}
	dp := make(map[int]int)
	for i := 0; i < n; i++ {
		if i != start1 {
			if nums[1][i] != 0 {
				dp[i] = nums[1][i]
			}
		}
	}
	final := true
	for final {

		final = true
	}
}
