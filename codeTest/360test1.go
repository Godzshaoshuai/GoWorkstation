package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	part := strings.Fields(line)
	n, _ := strconv.Atoi(part[0])
	k, _ := strconv.Atoi(part[1])
	line, _ = reader.ReadString('\n')
	part = strings.Fields(line)
	nums := make([]int, n+1)
	for i := 1; i <= n; i++ {
		nums[i], _ = strconv.Atoi(part[i-1])
	} //数据接受

	weight := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		weight[i] = make([]int, n+1)
	} //节点间最大权重共设置

	for i := 1; i < n+1; i++ {
		vmap := make(map[int]bool)
		sumW := 0
		for j := i; j < n+1; j++ {
			if !vmap[nums[j]] {
				vmap[nums[j]] = true
				sumW += nums[j]
			}
			weight[i][j] = sumW
		}
	} //权重设计完毕

	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, n)
		for j := 0; j <= k; j++ {
			dp[i][j] = -1
		}
	}
	dp[0][0] = 0 //DP初始化

	for i := 1; i < n+1; i++ {
		for j := 1; j <= min(i, k); j++ {
			for l := j - 1; l < i && l >= 0; l++ {
				if dp[l][j-1] != -1 {
					dp[i][j] = max(dp[i][j], dp[l][j-1]+weight[l+1][i])
				}

			}
		}
	}
	fmt.Println(dp[n][k])
}
