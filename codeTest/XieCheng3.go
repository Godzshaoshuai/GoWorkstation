package main

import "fmt"

func countk(n, k, x int) int {
	// jisuan qian n xulie weishu
	if n < x {
		return 0
	}
	return (n-x)/k + 1
}

func main() {
	var n, l, r, k, x int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&l, &r, &k, &x)
		count := countk(r, k, x) - countk(l-1, k, x)
		fmt.Println(count)
	}
}
