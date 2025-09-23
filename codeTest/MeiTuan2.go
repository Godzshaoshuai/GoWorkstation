package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	sort.Ints(nums)
	//fmt.Println(nums)
	dp := make([]int, n)
	nunsNew := []int{nums[0]}
	nunsNew2 := []int{}
	for i := 1; i < n; i++ {
		if nums[i] != nums[i-1] {
			nunsNew = append(nunsNew, nums[i])
		} else {
			nunsNew2 = append(nunsNew2, nums[i])
		}
	}
	for i := 0; i < len(nunsNew2); i++ {
		nunsNew = append(nunsNew, nunsNew2[i])
	}
	//fmt.Println(nunsNew)
	if nums[0] == 0 {
		dp[0] = 1
	} else {
		dp[0] = 0
	}
	total := dp[0]
	fmt.Fprintf(writer, "%d", nums[0])
	fmt.Fprintf(writer, " ")
	for i := 1; i < n; i++ {
		if nunsNew[i] == dp[i-1] {
			dp[i] = nunsNew[i] + 1
			//fmt.Println("+1", nums[i], dp[i])

		} else {
			dp[i] = dp[i-1]
			//fmt.Println("bu+1", nums[i], dp[i])
		}
		//fmt.Println(dp[i])
		total += dp[i]
		fmt.Fprintf(writer, "%d", nunsNew[i])
		if i < n-1 {
			fmt.Fprintf(writer, " ")
		}
	}
	fmt.Println(total)
	fmt.Fprintln(writer)

}
