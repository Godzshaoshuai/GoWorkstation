package main

import (
	"fmt"
)

func main() {
	a := 0
	fmt.Scan(&a)
	nums := make([]int, a)
	for i := 0; i < a; i++ {
		fmt.Scan(&nums[i])
	}
	//按位异或相等，暂定思路，nums整体异或，最稳定的方式就是所有的值完全按抑或取反，从最大的开始，但是这样肯定的k是缺失的
	//全体nums进行按位或操作，只要有零的地方，b数组的地点全零，如果有1的地方，未来b数组采取访问第一个a数组元素，使得b数组第一个值最小，以此确定k，然后后续顺延，所以就是找最小的k值，因为要零
	//k := 0
	//for i := 30; i >= 0; i-- {
	//	kNow := k | (1 << i)
	//	final := false
	//	for j := 0; j > a; j++ {
	//		if (nums[j] ^ kNow) < (nums[j] ^ k) { //genghuan
	//			final = true
	//			break
	//		} else if (nums[j] ^ kNow) > (nums[j] ^ k) { //zhongduan
	//			break
	//		}
	//	}
	//	if final {
	//		k = kNow
	//	}
	//}
	//for i := 0; i < a; i++ {
	//	fmt.Print(nums[i]^k, " ")
	//}
	if a == 0 {
		//6
		//1 1 4 5 1 4
	} else if a == 1 {
		fmt.Print(0)
	} else {
		numsb := make([]int, a)
		numsb[0] = 0
		k := numsb[0] ^ nums[1]
		for i := 1; i < a; i++ {
			numsb[i] = nums[i] ^ k
		}
		fmt.Println(numsb)
	}
}
