package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	T, _ := strconv.Atoi(scanner.Text())

	for t := 0; t < T; t++ {
		scanner.Scan()
		nm := strings.Split(scanner.Text(), " ")
		N, _ := strconv.Atoi(nm[0])
		M, _ := strconv.Atoi(nm[1])

		// 初始化邻接矩阵
		adj := make([][]bool, N+1)
		for i := range adj {
			adj[i] = make([]bool, N+1)
		}

		// 读取边信息
		for i := 0; i < M; i++ {
			scanner.Scan()
			xy := strings.Split(scanner.Text(), " ")
			x, _ := strconv.Atoi(xy[0])
			y, _ := strconv.Atoi(xy[1])
			adj[x][y] = true
			adj[y][x] = true // 无向图，双向都要标记
		}

		// 初始化分组，0表示未分组
		group := make([]int, N+1)
		currentGroup := 1

		// 为每个未分组的节点分配分组
		for i := 1; i <= N; i++ {
			if group[i] == 0 {
				// 该节点的邻居集合
				neighbors := make([]int, 0)
				for j := 1; j <= N; j++ {
					if i != j && adj[i][j] {
						neighbors = append(neighbors, j)
					}
				}

				// 检查是否有邻居已分配到当前组
				hasNeighborInGroup := false
				for _, j := range neighbors {
					if group[j] == currentGroup {
						hasNeighborInGroup = true
						break
					}
				}

				if hasNeighborInGroup {
					// 如果有邻居在当前组，需要创建新组
					currentGroup++
				}

				// 将当前节点分配到当前组
				group[i] = currentGroup

				// 处理与当前节点不相邻的节点（应属于同一组）
				for j := 1; j <= N; j++ {
					if i != j && !adj[i][j] && group[j] == 0 {
						group[j] = currentGroup
					}
				}
			}
		}

		// 验证分组是否符合完全多部图的条件
		isValid := true
		for i := 1; i <= N; i++ {
			for j := i + 1; j <= N; j++ {
				if group[i] == group[j] {
					// 同一组的节点必须没有边
					if adj[i][j] {
						isValid = false
						break
					}
				} else {
					// 不同组的节点必须有边
					if !adj[i][j] {
						isValid = false
						break
					}
				}
			}
			if !isValid {
				break
			}
		}

		if isValid {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
