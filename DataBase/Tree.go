package main

import "fmt"

type TreeNode struct {
	val   string
	left  *TreeNode
	right *TreeNode
}

func (this *TreeNode) PreOrder() {
	if this == nil {
		return
	}
	fmt.Print(this.val, " ")
	this.left.PreOrder()
	this.right.PreOrder()
}

func (this *TreeNode) MidOrder() {
	if this == nil {
		return
	}

	this.left.MidOrder()
	fmt.Print(this.val, " ")
	this.right.MidOrder()
}

func (this *TreeNode) PostOrder() {
	if this == nil {
		return
	}

	this.left.PostOrder()
	this.right.PostOrder()
	fmt.Print(this.val, " ")
}

func (this *TreeNode) LayerOrder() {
	if this == nil {
		return
	}
	numsIndex := []*TreeNode{}

	numsIndex = append(numsIndex, this)
	for numsIndex != nil {
		numsIndexNew := []*TreeNode{}
		for _, c := range numsIndex {
			fmt.Print(c.val, " ")
			if c.left != nil {
				numsIndexNew = append(numsIndexNew, c.left)
			}
			if c.right != nil {
				numsIndexNew = append(numsIndexNew, c.right)
			}
		}
		numsIndex = numsIndexNew
	}
}

func main() {
	t := &TreeNode{val: "A"}
	t.left = &TreeNode{val: "B"}
	t.right = &TreeNode{val: "C"}
	t.left.left = &TreeNode{val: "D"}
	t.left.right = &TreeNode{val: "E"}
	t.right.left = &TreeNode{val: "F"}
	fmt.Println("先序排序：")
	t.PreOrder()
	fmt.Println("\n中序排序：")
	t.MidOrder()
	fmt.Println("\n后序排序")
	t.PostOrder()
	fmt.Println("\n层次便利")
	t.LayerOrder()
}
