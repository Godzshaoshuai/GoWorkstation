package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, a, b int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &n)
	for p := 0; p < n; p++ {
		fmt.Fscanf(reader, "%d %d\n", &a, &b)
		nums := make([][]int, a)
		dp := make([][]int, a)
		lujing := make([][]int, a)
		for i := 0; i < a; i++ {
			nums[i] = make([]int, b)
			dp[i] = make([]int, b)
			lujing[i] = make([]int, b)

			for j := 0; j < b; j++ {
				fmt.Fscanf(reader, "%d", &nums[i][j])
				//fmt.Println(nums[i][j])
			}
			fmt.Fscanf(reader, "\n")

		}
		dp[0][0] = nums[0][0]
		lujing[0][0] = 0
		for i := 0; i < a; i++ {
			for j := 0; j < b; j++ {
				if j >= 1 && i >= 1 {
					dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + nums[i][j]
					if dp[i][j] == dp[i-1][j]+nums[i][j] {
						lujing[i][j] = 1
					} else {
						lujing[i][j] = -1
					}
				} else if i == 0 && j >= 1 {
					dp[i][j] = dp[i][j-1] + nums[i][j]
					lujing[i][j] = -1
				} else if i >= 0 && j == 1 {
					dp[i][j] = dp[i-1][j] + nums[i][j]
					lujing[i][j] = 1
				} else {
				}
			}
		}
		minX := dp[a-1][b-1]
		for i := a - 1; i >= 0; {
			for j := b - 1; j >= 0; {
				if minX > dp[i][j] {
					minX = dp[i][j]
					//fmt.Println(minX, dp[i][j])
				}
				if lujing[i][j] == 0 {
					i = -999
					j = -999
				} else if lujing[i][j] == 1 {
					i--
				} else if lujing[i][j] == -1 {
					j--
				}
			}
		}
		//fmt.Println(dp)
		//fmt.Println(nums)
		//fmt.Println(lujing)
		if minX >= 0 {
			fmt.Println(1)
		} else {
			fmt.Println(1 - minX)
		}
	}
}
