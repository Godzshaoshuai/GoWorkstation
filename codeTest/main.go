package main

import "fmt"

func compareABC(nums1, nums2 []int) int {
	if nums1[0] > nums2[0] && nums1[1] > nums2[1] && nums1[2] > nums2[2] {
		return 1
	}
	if nums1[0] < nums2[0] && nums1[1] < nums2[1] && nums1[2] < nums2[2] {
		return -1
	}

	return 0
}

func main() {
	var t, a, b, c int
	fmt.Scanln(&t)
	nums := make([][]int, t)
	for i := 0; i < t; i++ {
		fmt.Scanln(&a, &b, &c)
		nums[i] = []int{a, b, c}
	}
	visit := make([]bool, t)
	count := 0
outerLoop:
	for i := 0; i < t; i++ {
		if !visit[i] {
			for j := i + 1; j < t; j++ {
				if !visit[j] {
					test := compareABC(nums[i], nums[j])
					if test == 1 {
						count++
						visit[j] = true
					} else if test == -1 {
						count++
						visit[i] = true
						if i >= t {
							break outerLoop
						}
						continue outerLoop
					}
				}

			}
		}

	}
	fmt.Println(count)
}
