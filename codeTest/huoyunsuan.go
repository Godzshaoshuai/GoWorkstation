package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func baohan(a, b int) int {
	if a|b == a {
		return a
	}
	if a|b == b {
		return b
	}
	return a + b
}

func main() {
	var t int
	fmt.Scanln(&t)
	nums := make([]int, t)
	list := make([]int, 0)
	inputs := bufio.NewScanner(os.Stdin)
	inputs.Scan()                             //每次读入一行
	data := strings.Split(inputs.Text(), " ") //通过空格将他们分割，并存入一个字符串切片
	for i := range data {
		nums[i], _ = strconv.Atoi(data[i]) //将字符串转换为int
	}
	for i := 0; i < t; i++ {
		visit := true
		listVisit := make([]int, 0)
		for j := 0; j < len(list); j++ {
			if baohan(list[j], nums[i]) == list[j] {
				visit = false
				break
			}
			if baohan(list[j], nums[i]) == nums[i] {
				listVisit = append(listVisit, j)
			}
		}
		if visit {
			list = append(list, nums[i])
			if len(listVisit) >= 1 {
				for j := len(listVisit) - 1; j >= 0; j-- {
					list = append(list[:listVisit[j]], list[listVisit[j]+1:]...)
				}
			}

		}
	}
	fmt.Println(len(list))
}
