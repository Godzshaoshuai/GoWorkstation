package FastCode

// 给定数组和目标，每个数之前任取+-符号，最后达成目标的个数
func findTargetSumWays(nums []int, target int) int {
	sum := sum(nums)
	n := len(nums)
	neg := sum - target
	if neg < 0 || neg%2 == 1 {
		return 0
	}
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, neg+1)
	}
	dp[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j <= neg/2; j++ {
			dp[i+1][j] = dp[i][j]
			if j >= nums[i] {
				dp[i+1][j] += dp[i][j-nums[i]]
			}
		}
	}
	return dp[n][neg/2]
}

func sum(nums []int) int {
	sum := 0
	for _, c := range nums {
		sum += c
	}

	return sum
}
