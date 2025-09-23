package main

import "fmt"

type LinkNode struct {
	val      int
	NextNode *LinkNode
}

// 循环链表
type Ring struct {
	val        interface{}
	next, prev *Ring
}

func (this *Ring) initRing() *Ring {
	this.prev = nil
	this.next = nil
	return this
}

// 创建长度为n的循环链表
func nRing(n int) *Ring {
	if n <= 0 {
		return nil
	}
	r := new(Ring)

	p := r
	for i := 0; i < n; i++ {
		p.next = &Ring{prev: p}
		p = p.next
	}
	p.next = r
	r.prev = p
	return r
}

func main() {
	node := new(LinkNode)
	node.val = 2
	fmt.Println(node)
	node1 := LinkNode{
		val:      1,
		NextNode: nil,
	}
	fmt.Println(node1)
	node2 := &LinkNode{
		val:      2,
		NextNode: nil,
	}
	fmt.Println(node2)

}
