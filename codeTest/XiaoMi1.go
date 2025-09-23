package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, p, q int
	fmt.Fscanf(reader, "%d %d %d\n", &n, &p, &q)
	leftShoes := make([]map[int]int, 50)
	rightShoes := make([]map[int]int, 50)
	for i := 0; i < 50; i++ {
		leftShoes[i] = make(map[int]int)
		rightShoes[i] = make(map[int]int)
	}
	for i := 0; i < n; i++ {
		var a, b, c int
		fmt.Fscanf(reader, "%d %d %d\n", &a, &b, &c) //信息输入，后续左右脚录入
		if a == 0 {
			leftShoes[b][c]++
		} else {
			rightShoes[b][c]++
		}
	}

	var finalMoney int64 = 0 //最终最大金额记录，这个本质上就是一个贪心计算，并且限制了不同色价格低于同色，并且不同size无法输出

	for sizeIndex := 35; sizeIndex < 49; sizeIndex++ {
		leftCount := leftShoes[sizeIndex]
		rightCount := rightShoes[sizeIndex]
		if len(leftCount) == 0 || len(rightCount) == 0 {
			continue
		}
		countDoubleNew := 0
		for k, v := range leftCount {
			if numbxx, ok := rightCount[k]; ok {
				paris := min(v, numbxx)
				countDoubleNew += paris
			}
		}
		countl := 0
		for _, count := range leftCount {
			countl += count
		}
		countr := 0
		for _, count := range rightCount {
			countr += count
		}
		maxNumb := min(countl, countr) //最大怕配对数量
		notPair := maxNumb - countDoubleNew
		finalMoney += int64(countDoubleNew*p + notPair*q)
	}
	fmt.Println(finalMoney)
}
