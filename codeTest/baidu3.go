package main

import "fmt"

func main() {
	var n, m, l, r, x int
	fmt.Scan(&n, &m)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])

	}
	for i := 0; i < m; i++ {
		fmt.Scan(&l, &r, &x)
		k := nums[l-1] | x
		//fmt.Println(nums, x, k)
		j := l
		for ; j < r; j++ {
			if nums[j]|x != k {
				break
			}
		}
		if j == r && nums[j-1]|x == k {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
