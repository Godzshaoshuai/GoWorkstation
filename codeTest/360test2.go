package main

import (
	"fmt"
)

type UnionFind struct {
	parent []int
	size   []int
} //并查集

func NewUnionFind(n int) *UnionFind {
	u := &UnionFind{
		parent: make([]int, n),
		size:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		u.parent[i] = i
		u.size[i] = 1
	}
	return u
}

func (this *UnionFind) Find(num int) int {
	if this.parent[num] != num {
		this.parent[num] = this.Find(this.parent[num])
	}
	return this.parent[num]
}

func (this *UnionFind) Union(a, b int) int64 {
	parentA := this.Find(a)
	parentB := this.Find(b)
	if parentA == parentB {
		return 0
	}
	sizeA := int64(this.size[parentA])
	sizeB := int64(this.size[parentB])
	if this.size[parentA] < this.size[parentB] {
		this.parent[parentA] = parentB
		this.size[parentB] += this.size[parentA]
	} else {
		this.parent[parentB] = parentA
		this.size[parentA] += this.size[parentB]
	}

	return int64(2 * sizeB * sizeA)
}

func main() {
	var n int
	var a, b int64
	fmt.Scan(&n)
	position := make(map[int64]int) //强制寻找位置
	var sumR int64 = 0
	offser := [][2]int{{0, 1}, {0, -1}, {1, 0}, {1 - 1}, {1, 1}, {-1, 0}, {-1, -1}, {-1, 1}}
	UnionE := NewUnionFind(n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a, &b)
		key := a*1000000000*2 + 3 + b //区分不开就2运算吧
		position[key] = i
		sumR += 1
		for j := 0; j < 8; j++ {
			newA := a + int64(offser[j][0])
			newB := b + int64(offser[j][1])
			newKey := newA*1000000000*2 + 3 + newB
			if idnex, ok := position[newKey]; ok {
				plus := UnionE.Union(i, idnex)
				sumR += plus
			} //如果存在，合并，并修改插入值
		}
		fmt.Println(sumR)
	}

}
