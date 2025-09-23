package main

import (
	"fmt"
	"sync"
)

type DoubleList struct {
	head *LinkNodeDouble
	tail *LinkNodeDouble
	len  int
	lock sync.Mutex
}

type LinkNodeDouble struct {
	next, pre *LinkNodeDouble
	val       string
}

func (this *LinkNodeDouble) GetValue() string {
	return this.val
}

func (this *LinkNodeDouble) GetPre() *LinkNodeDouble {
	return this.pre
}

func (this *LinkNodeDouble) GetNext() *LinkNodeDouble {
	return this.next
}

func (this *LinkNodeDouble) HashNext() bool {
	return this.next != nil
}

func (this *LinkNodeDouble) HashPre() bool {
	return this.pre != nil
}

func (this *LinkNodeDouble) IsEmpty() bool {
	return this == nil
}

func (this *DoubleList) GetLen() int {
	return this.len
}

func initDouble() *DoubleList {
	headNew := new(LinkNodeDouble)
	tailNew := new(LinkNodeDouble)
	headNew.next = tailNew
	headNew.pre = nil
	tailNew.pre = headNew
	tailNew.next = nil
	return &DoubleList{
		head: headNew,
		tail: tailNew,
		len:  0,
	}
}

func (this *DoubleList) PushHead(v string) {
	this.lock.Lock()
	defer this.lock.Unlock()

	this.len++
	node := new(LinkNodeDouble)
	node.val = v
	node.next = this.head.next
	node.pre = this.head
	this.head.next.pre = node
	this.head.next = node
}

func (this *DoubleList) PopHead() string {
	this.lock.Lock()
	defer this.lock.Unlock()
	if this.len == 0 {
		panic("empty")
	}
	this.len--
	s := this.head.next.val
	this.head.next.next.pre = this.head.next
	this.head.next = this.head.next.next
	return s
}

func (this *DoubleList) GetHead() string {
	return this.head.next.val
}

func (this *DoubleList) PushTail(v string) {
	this.lock.Lock()
	defer this.lock.Unlock()

	this.len++
	node := new(LinkNodeDouble)
	node.val = v
	this.tail.pre.next = node
	node.pre = this.tail.pre
	this.tail.pre = node
	node.next = this.tail
}

func (this *DoubleList) PopTail() string {
	this.lock.Lock()
	defer this.lock.Unlock()

	if this.len == 0 {
		panic("empty")
	}
	this.len--
	s := this.tail.pre.val
	this.tail.pre.pre.next = this.tail
	this.tail.pre = this.tail.pre.pre
	return s
}

func (this *DoubleList) GetTail() string {
	return this.tail.pre.val
}

func main() {
	doubleList := initDouble()
	doubleList.PushHead("one")
	doubleList.PushTail("two")
	fmt.Println(doubleList.GetHead(), doubleList.GetTail(), doubleList.GetLen())
	fmt.Println("head string,", doubleList.PopHead())
	fmt.Println("head string2,", doubleList.PopHead())
	doubleList.PushHead("one")
	fmt.Println("tail string,", doubleList.PopTail())
	fmt.Println("tail string for empty,", doubleList.PopTail())
}
