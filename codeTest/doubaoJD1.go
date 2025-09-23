package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], _ = strconv.Atoi(parts[i])
	}

	// 计算每个位置的LIS长度和计数
	dpLen, dpCnt := computeDp(a, n)

	// 计算每个位置的反向LIS长度和计数
	revA := make([]int, n)
	for i := 0; i < n; i++ {
		revA[i] = a[n-1-i]
	}
	rdpLen, rdpCnt := computeDp(revA, n)

	// 反转反向计算的结果以匹配原始数组
	reverseDp(rdpLen, rdpCnt, n)

	// 找出LIS的最大长度
	maxLen := 0
	for _, l := range dpLen {
		if l > maxLen {
			maxLen = l
		}
	}

	// 计算总共有多少个LIS
	total := 0
	for i := 0; i < n; i++ {
		if dpLen[i] == maxLen {
			total += dpCnt[i]
		}
	}

	// 计算每个元素在多少个LIS中出现
	result := make([]int, n)
	for i := 0; i < n; i++ {
		if dpLen[i]+rdpLen[i]-1 == maxLen {
			result[i] = (dpCnt[i] * rdpCnt[i] * total) / (dpCnt[i] * rdpCnt[i])
		} else {
			result[i] = 0
		}
	}

	// 输出结果
	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(result[i])
	}
	fmt.Println()
}

// 计算每个位置的LIS长度和计数
func computeDp(a []int, n int) ([]int, []int) {
	dpLen := make([]int, n)
	dpCnt := make([]int, n)
	tails := make([]int, 0)    // 存储各长度LIS的最小尾元素
	counts := make([][]int, 0) // 存储各长度LIS的计数信息

	for i := 0; i < n; i++ {
		x := a[i]

		// 找到第一个大于等于x的位置
		idx := sort.Search(len(tails), func(j int) bool {
			return tails[j] >= x
		})

		// 计算当前位置的LIS长度
		if idx == 0 {
			dpLen[i] = 1
			dpCnt[i] = 1
		} else {
			dpLen[i] = idx + 1
			// 累加前一个长度所有小于x的计数
			dpCnt[i] = 0
			for _, p := range counts[idx-1] {
				if p < x {
					dpCnt[i] += p >> 16
				}
			}
		}

		// 更新tails和counts数组
		if idx == len(tails) {
			tails = append(tails, x)
			counts = append(counts, []int{(x << 16) | dpCnt[i]})
		} else {
			if x < tails[idx] {
				tails[idx] = x
				counts[idx] = []int{(x << 16) | dpCnt[i]}
			} else if x == tails[idx] {
				counts[idx] = append(counts[idx], (x<<16)|dpCnt[i])
			}
		}
	}

	return dpLen, dpCnt
}

// 反转反向计算的结果以匹配原始数组
func reverseDp(rdpLen, rdpCnt []int, n int) {
	for i := 0; i < n/2; i++ {
		j := n - 1 - i
		rdpLen[i], rdpLen[j] = rdpLen[j], rdpLen[i]
		rdpCnt[i], rdpCnt[j] = rdpCnt[j], rdpCnt[i]
	}
}
