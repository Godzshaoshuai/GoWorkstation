package main

import (
	"fmt"
)

const InfMax = 999999

func Dijkstra(graph [][]int, startIndex int) []int {
	n := len(graph)
	if n == 0 || startIndex < 0 || startIndex >= n {
		return nil // 无效的输入
	}
	finalNums := graph[startIndex]
	visit := make([]bool, n)
	visit[startIndex] = true
	for i := 0; i < n-1; i++ {
		minIndex := InfMax
		minNum := InfMax
		for j := 0; j < n; j++ {
			if !visit[j] && finalNums[j] < minNum {
				minNum = finalNums[j]
				minIndex = j
			}
		}
		if minNum != InfMax {
			visit[minIndex] = true
			for j := 0; j < n; j++ {
				if graph[minIndex][j]+finalNums[minIndex] < finalNums[j] {
					finalNums[j] = graph[minIndex][j] + finalNums[minIndex]
				}
			}
		}
	}
	for i, _ := range finalNums {
		if finalNums[i] == InfMax {
			finalNums[i] = -1
		}

	}
	return finalNums
}

func main() {
	// 测试用例1: 连通图
	fmt.Println("===== 测试用例1: 连通图 =====")
	graph1 := [][]int{
		{0, 4, InfMax, 2, InfMax}, // 节点0
		{4, 0, 1, 5, InfMax},      // 节点1
		{InfMax, 1, 0, InfMax, 3}, // 节点2
		{2, 5, InfMax, 0, 1},      // 节点3
		{InfMax, InfMax, 3, 1, 0}, // 节点4
	}
	distances1 := Dijkstra(graph1, 0)
	fmt.Println(distances1)

	// 测试用例2: 包含不可达节点的图
	fmt.Println("\n===== 测试用例2: 包含不可达节点 =====")
	graph2 := [][]int{
		{0, 2, 3, InfMax, InfMax},           // 节点0
		{2, 0, 1, 5, InfMax},                // 节点1
		{3, 1, 0, InfMax, InfMax},           // 节点2
		{InfMax, 5, InfMax, 0, InfMax},      // 节点3
		{InfMax, InfMax, InfMax, InfMax, 0}, // 节点4（与其他节点完全隔离）
	}
	distances2 := Dijkstra(graph2, 0)
	fmt.Println(distances2)

	// 测试用例3: 单节点图
	fmt.Println("\n===== 测试用例3: 单节点图 =====")
	graph3 := [][]int{{0}} // 只有节点0
	distances3 := Dijkstra(graph3, 0)
	fmt.Println(distances3)
	index1 := 0
	index2 := 0
	value := 0
	n := 0
	m := 0
	fmt.Println("Put in n and m")
	fmt.Scanln(&n, &m)
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
		for j := 0; j < n; j++ {
			graph[i][j] = InfMax
		}
	}
	//fmt.Scanln(&m)
	fmt.Println("Put in index1,index2,Value")
	for i := 0; i < m; i++ {
		fmt.Scanln(&index1, &index2, &value)
		graph[index1][index2] = value
	}
	distances4 := Dijkstra(graph, 0)
	fmt.Println(distances4)
}
