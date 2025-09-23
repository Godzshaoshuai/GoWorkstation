package main

import (
	"fmt"
	"sort"
)

// 计算数字的各位平方和
func digitSquareSum(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}
	return sum
}

// 优化解法：直接构造最小的关联数
func findMinAssociatedNumber(n int) int {
	target := digitSquareSum(n)

	// 特殊情况处理
	if target == 0 {
		return 0 // 按题意不会出现，因为n是正整数
	}

	// 从最大可能的位数开始尝试
	for length := 1; ; length++ {
		// 检查当前长度是否可能
		minSum := length * 1  // 最小平方和(全是1)
		maxSum := length * 81 // 最大平方和(全是9)

		if target < minSum || target > maxSum {
			continue
		}

		// 尝试构造当前长度的数
		digits := make([]int, 0, length)
		remaining := target

		// 从大到小选择数字，尽可能使用大数字
		for i := 0; i < length; i++ {
			// 最后一位特殊处理，确保总和刚好等于target
			if i == length-1 {
				if remaining >= 1 && remaining <= 81 {
					// 找到对应的数字(1-9)
					d := 1
					for d*d < remaining {
						d++
					}
					if d*d == remaining {
						digits = append(digits, d)
						remaining = 0
						break
					}
				}
			} else {
				// 选择最大可能的数字，但要保证剩余位数能达到剩余的和
				for d := 9; d >= 1; d-- {
					square := d * d
					remainingDigits := length - i - 1
					minRemaining := remainingDigits * 1
					maxRemaining := remainingDigits * 81

					if square <= remaining && (remaining-square) >= minRemaining && (remaining-square) <= maxRemaining {
						digits = append(digits, d)
						remaining -= square
						break
					}
				}
			}
		}

		// 如果成功构造出符合条件的数字组合
		if remaining == 0 && len(digits) == length {
			// 排序使数字最小
			sort.Ints(digits)

			// 转换为整数
			result := 0
			for _, d := range digits {
				result = result*10 + d
			}
			return result
		}
	}
}

func main() {
	var n int
	fmt.Print("请输入整数n: ")
	fmt.Scan(&n)

	result := findMinAssociatedNumber(n)
	fmt.Printf("与%d关联且每位都不为0的最小正整数是: %d\n", n, result)

	// 验证结果
	fmt.Printf("验证: %d的各位数字平方和为%d, %d的各位数字平方和为%d\n",
		n, digitSquareSum(n),
		result, digitSquareSum(result))
}
