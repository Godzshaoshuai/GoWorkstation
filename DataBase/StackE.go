package main

import (
	"fmt"
	"sync"
)

type ArrayStack struct {
	array []string
	size  int
	lock  sync.Mutex
}

func (strak *ArrayStack) Push(v string) {
	strak.lock.Lock()
	defer strak.lock.Unlock()

	strak.array = append(strak.array, v)
	strak.size++
}

func (stack *ArrayStack) Pop() string {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	if stack.size == 0 {
		panic("empty")
	}

	s := stack.array[stack.size-1]
	stack.size--
	stack.array = stack.array[:stack.size]

	return s
}

func (stcak *ArrayStack) get() string {
	if stcak.size == 0 {
		panic("empty for get")
	}
	return stcak.array[stcak.size-1]
}

func main() {
	arrayStack := new(ArrayStack)
	arrayStack.Push("First")
	fmt.Println(arrayStack.array, arrayStack.size)
	arrayStack.Push("Second")
	fmt.Println(arrayStack.array, arrayStack.size)
	arrayStack.Push("Third")
	fmt.Println(arrayStack.array, arrayStack.size)
	fmt.Println("pop", arrayStack.Pop(), arrayStack.size)
	fmt.Println(arrayStack.array, arrayStack.size, arrayStack.get())
	fmt.Println("pop", arrayStack.Pop(), arrayStack.size)
	fmt.Println("pop", arrayStack.Pop(), arrayStack.size)
	fmt.Println("pop", arrayStack.Pop())

}
