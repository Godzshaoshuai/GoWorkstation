package main

import (
	"fmt"
)

//func dfs_red(redNums map[int]bool, D [][]int, n int) []int {
//
//}

func main() {
	var n, m, k, a, b, c int //这个题限制了有n个节点，并且有n-1行的边权重筛选，并且没有指定不可达，默认可达，还说是树结构，其实本质上就是一个无环顺序链表
	fmt.Scan(&n, &m, &k)

	redmap := make(map[int]bool)
	dp1 := make([]int, n) //qianxu
	dp2 := make([]int, n) //houxu
	for i := 0; i < m; i++ {
		fmt.Scan(&a)
		redmap[a] = true //shanchu
	}
	index := make(map[int]int, n)
	index2 := make(map[int]int, n)
	value := make([]int, n)
	for i := 0; i < n-1; i++ {
		fmt.Scan(&a, &b, &c)

		if a < b {
			value[a-1] = c
			index[a] = b
			index2[b] = a

		} else {
			value[b-1] = c
			index[b] = a
			index2[a] = b
		}
	}
	if redmap[1] {
		dp1[0] = 0
	} else {
		dp1[0] = 99999
	}
	for i := index[1]; i <= n; {
		if _, ok := index[i]; !ok {
			break
		}
		if redmap[i] {
			dp1[i-1] = 0
		} else {
			dp1[i-1] = dp1[index2[i]-1] + value[index2[i]-1]
		}
		i = index[i]
	}
	if redmap[n] {
		dp2[n-1] = 0
	} else {
		dp2[n-1] = 99999
	}
	for i := index2[n]; i <= n; {
		if _, ok := index2[i]; !ok {
			break
		}
		if redmap[i] {
			dp2[i-1] = 0
		} else {
			dp2[i-1] = dp2[index[i]-1] + value[i-1]
		}
		i = index2[i]
	}
	for i := 0; i < n; i++ {
		if dp1[i] > dp2[i] {
			dp1[i] = dp2[i]
		}
		fmt.Print(dp1[i])
		if i < n-1 {
			fmt.Print(" ")
		}
	}
	if k == 1 {
		for i := 0; i < n; i++ {
			fmt.Print(dp1[i])
			if i < n-1 {
				fmt.Print(" ")
			}
		}
	} else {
		for i := 0; i < n; i++ {
			valuex := 0
			for j := 0; j < k; j++ {
				valuex += dp1[i+j]
			}
		}
	}

}
