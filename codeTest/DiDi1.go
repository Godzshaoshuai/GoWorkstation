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
	n, _ := strconv.Atoi(strings.TrimSpace(line))
	powerAndCost := make([][2]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &powerAndCost[i][0], &powerAndCost[i][1])
	}
	const maxPower = 6000
	const inf = 9000000
	const maxIndex = 2*maxPower + 1
	dp := make([]int, maxIndex) //以最大值6000为0值，防止数组越界
	for i := range dp {
		dp[i] = inf
	}
	dp[maxPower] = 0
	for i := 0; i < n; i++ {
		poweri := powerAndCost[i][0]
		costi := powerAndCost[i][1]
		if poweri >= 0 {
			for j := maxIndex - 1 - poweri; j >= 0; j-- {
				if dp[j] != inf {
					if dp[j+poweri] > dp[j]+costi {
						dp[j+poweri] = dp[j] + costi
					}
				}
			}
		} else {
			for j := -poweri; j < maxIndex; j++ {
				if dp[j] != inf {
					if dp[j+poweri] > dp[j]+costi {
						dp[j+poweri] = dp[j] + costi
					}
				}
			}
		}
	}
	if dp[maxPower+1] != inf {
		fmt.Println(dp[maxPower+1])
	} else {
		fmt.Println(-1)
	}
}
