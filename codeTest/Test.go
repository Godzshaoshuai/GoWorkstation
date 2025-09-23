package main

import (
	"bufio"
	"fmt"
	"os"
)

type node1 struct {
	val  int
	next *node1
}

func initnode1() *node1 {
	return &node1{
		val:  9999,
		next: nil,
	}
}

func (this *node1) insertnode1(x int) {
	temp := this
	if temp.next != nil && x > temp.next.val {
		for ; temp.next != nil; temp = temp.next {
			if x > temp.val && x < temp.next.val {
				break
			}
		}
	}
	node1New := initnode1()
	node1New.val = x
	node1New.next = temp.next
	temp.next = node1New
}

func (this *node1) delenode1() {
	if this.next != nil {
		this.next = this.next.next
	}
}

func (this *node1) query() {
	if this.next != nil {
		fmt.Println(this.next.val)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, m, x int
	fmt.Fscan(reader, &n)
	head := initnode1()
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &m)

		switch m {
		case 1:
			{
				fmt.Fscan(reader, &x)
				head.insertnode1(x)
			}
		case 2:
			{
				head.query()
			}
		case 3:
			{
				head.delenode1()
			}
		default:
			{
			}
		}
	}
}
