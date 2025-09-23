package main

import "fmt"

// 冒泡排序：
// 遍历n-1次，每次如果大就交换位置，以此更迭
// 最佳 O(N) 最差 O(n2) 平均 O(n2) 稳定
func BubbleSort(list []int) {
	n := len(list)
	for i := n - 1; i > 0; i-- {
		noSwap := false //若不存在交换，快速截止标识
		for j := 0; j < i; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				noSwap = true
			}
		}
		if !noSwap {
			return
		}
	}
}

// 选择排序：
// 每次选择一个最，然后与还没排序的地方交换，但是极度不稳定，如果数组里面存在相同数值
// 最佳 O(N2) 最差 O(n2) 平均 O(n2) 不稳定
func SelectSort(list []int) {
	n := len(list)
	for i := 0; i < n; i++ {
		index := i
		for j := i + 1; j < n; j++ {
			if list[j] < list[index] {
				index = j
			}
		}
		list[i], list[index] = list[index], list[i]
	}
}

// 插入排序：
// 一个逐渐丰富数组的排序，也就是便利数组，然后看下一个数应该在已经排序的地方的哪里，以此进行插入排序
// 最佳 O(N) 最差 O(n2) 平均 O(n2) 稳定
func InsertSprt(list []int) {
	n := len(list)
	for i := 1; i < n; i++ {
		tempInt := list[i]
		j := i - 1
		if tempInt < list[j] {
			for ; j >= 0 && tempInt < list[j]; j-- {
				list[j+1] = list[j]
			}
			list[j+1] = tempInt
		}

	}
}

// 希尔排序：
// 希尔排序是直接插入排序的升级版， 因为直接插入排序的最优时刻是几乎已经完全实现排序，所以希尔排序就按照步长先进行间隔的排序，最终随着步长的缩短，最后就是完全排序了
// 起始设置步长是长度的一半，这样对比后，左侧就是小的数，右侧就是大的数，后续数组可以认为数一个泛化的左小右大，但是并不是左边的最大值也小于右边，那是快速排序，然后step再÷2
// 虽然不能保证一次一定会遇到之前的数，但是数据分布更均衡了，之后一直到步长是1，进行一遍排序即可，其实预期说是插入排序优化，我更认为是局部冒泡排序

func ShellSort(list []int) {
	n := len(list)
	for step := n / 2; step >= 1; step /= 2 {
		for i := step; i < n; i += step {
			for j := i - step; j > 0; j -= step {
				if list[j+step] < list[j] {
					list[j], list[j+step] = list[j+step], list[j]
					continue
				}
				break
			}
		}
	}

}

// 归并排序：
// 非常高效的排序方式，采取的左右分层排序，但是由于调用了递归，所以实际是超细致的局部排序，然后合并，符合算法的整体优化原则
// 最佳 O(NlogN) 最差 O(NlogN) 平均 O(NlogN) 稳定
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// 堆排序：
// 同样是不稳定的算法，无法保证相同的值的前后相对方向位置不变。整体思路是将数组视作一个完全二叉树的的下标情况，受限进行一个初始化，之后逐个交替堆排序
// 最佳 O(NlogN) 最差 O(NlogN) 平均 O(NlogN) 不稳定
func Heapbuild(list []int, n, index int) {
	largestIndex := index
	left := 2*index + 1
	right := 2*index + 2

	if left < n && list[left] > list[largestIndex] {
		largestIndex = left
	}
	if right < n && list[right] > list[largestIndex] {
		largestIndex = right
	}

	if largestIndex != index {
		list[largestIndex], list[index] = list[index], list[largestIndex]
		Heapbuild(list, n, largestIndex)
	}
}

func HeapSort(list []int) {
	n := len(list)
	for i := n/2 - 1; i >= 0; i-- {
		Heapbuild(list, n, i)
	}
	for i := n - 1; i > 0; i-- {
		list[0], list[i] = list[i], list[0]
		Heapbuild(list, i, 0)
	}
}

// 快速排序：
// 类似于归并排序，归并排序的递归主要是拆然后和，快速的拆法是根据最后元素对对应区域的元素进行划分，实现小的在左，大的在右，逐步实现，知道最终区间step无差异
// 最佳 O(NlogN) 最差 O(NlogN) 平均 O(N2) 不稳定
func QuickSort(list []int, begin, end int) {
	if begin < end {
		pivot := partition(list, begin, end)
		QuickSort(list, begin, pivot-1)
		QuickSort(list, pivot+1, end)
	}

}

func partition(list []int, begin, end int) int {
	pivot := list[end]
	i := begin - 1
	for j := begin; j < end; j++ {
		if list[j] <= pivot {
			i++
			list[i], list[j] = list[j], list[i]
		}
	}
	list[i+1], list[end] = list[end], list[i+1]
	return i + 1
}

func main() {
	list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	//BubbleSort(list)
	//SelectSort(list)
	//InsertSprt(list)
	//SelectSort(list)
	//list = MergeSort(list)
	//HeapSort(list)
	QuickSort(list, 0, len(list)-1)
	fmt.Println(list)
}
