package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

//func insertNumIndex(nums []int, value int) int {
//	for i := 0; i < len(nums)-1; i++ {
//		if nums[i] < value && nums[i+1] > value {
//			return i
//		}
//	}
//	return len(nums) - 1
//}

func main() {
	var T, n, m int
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	T, _ = strconv.Atoi(strings.TrimSpace(line))
	for p := 0; p < T; p++ {
		line, _ = reader.ReadString('\n')
		parts := strings.Fields(line)
		n, _ = strconv.Atoi(parts[0])
		m, _ = strconv.Atoi(parts[1])
		//fmt.Fscan(reader, &n, &m)
		nums := make([]int, n)
		total := 0
		line, _ = reader.ReadString('\n')
		parts = strings.Fields(line)
		for o := 0; o < n; o++ {
			nums[o], _ = strconv.Atoi(parts[o])
			total += nums[o]
		}
		sort.Ints(nums)
		averageNums := total / n
		yushuNums := total % n
		minSteps := 0
		for i := n - 1; i >= 0; i-- {
			if nums[i] > averageNums {
				minSteps += nums[i] - averageNums
			} else {
				break
			}
		}
		if m >= minSteps {
			if yushuNums == 0 {
				fmt.Println(0)
			} else {
				fmt.Println(1)
			}
		} else {
			for i := 0; i < m; i++ {
				max1 := nums[n-1]
				min1 := nums[0]
				if max1 == min1 {
					break
				}
				nextMaxIndex := n - 1
				nextMinIndex := 0
				for j := n - 1; j > 0; j-- {
					if nums[j] < max1 {
						nextMaxIndex = j
						break
					}
				}
				for j := 0; j < n; j++ {
					if nums[j] > min1 {
						nextMinIndex = j
						break
					}
				}
				nums[n-1]--
				nums[0]++
				nums[n-1], nums[nextMaxIndex+1] = nums[nextMaxIndex+1], nums[n-1]
				nums[0], nums[nextMinIndex-1] = nums[nextMinIndex-1], nums[0]
			}
			fmt.Println(nums[n-1] - nums[0])
		}

	}
}
