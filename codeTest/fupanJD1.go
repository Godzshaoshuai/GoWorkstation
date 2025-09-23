package main

import "fmt"

type node struct {
	size int
	val  int
	next *node
}

func main() {
	var n, a int
	fmt.Println("shuru")
	fmt.Scanln(&n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&a)
		nums[i] = a
	}
	nodes := make([]*node, 0)
	max := 1
	count := make(map[int]int)
	for i := 0; i < n; i++ {
		count[nums[i]] = 0
		countx := 0
		for _, c := range nodes {
			fmt.Println("cval", c.val, "nums", nums[i], "nodes长度", len(nodes))
			if c.val < nums[i] { //先便利大小，如果有比初始值大的就进入找尾部，如果大于最后的尾部，那就插入，否则就复制一个新的
				fmt.Println("大小比较进入", countx)
				countx++

				head := c
				for head.next != nil {
					head = head.next
					fmt.Print("当前遍历值", head.val)
				}
				fmt.Println()
				if nums[i] > head.val { //比尾部都大，直接插入
					nodeNew := &node{
						val:  nums[i],
						size: 1,
						next: nil,
					}
					head.next = nodeNew
					c.size = c.size + 1
					if max < c.size {
						max = c.size
					}
					fmt.Println("尾部插入")
				} else { //没有尾部大，创建新的c，插入后进入noeds
					nodeNew := &node{ //初始化头节点
						val:  c.val,
						size: 1,
						next: nil,
					}
					temp := nodeNew
					head = c.next
					for ; head != nil; head = head.next {
						nodeNew.size++
						if head.val < nums[i] {
							fmt.Print(head.val)
							temp.next = head
							temp = temp.next
						} else {
							temp.next = &node{
								val:  nums[i],
								size: nodeNew.size,
								next: nil,
							}
							break
						}
					}
					if max < nodeNew.size {
						max = nodeNew.size
					}
					fmt.Println("中间插入")
					nodes = append(nodes, nodeNew)
				}
			}
		}
		if countx == 0 { //没有插入，需要在初始nods加入该节点
			nodeNew := &node{
				val:  nums[i],
				size: 1,
				next: nil,
			}
			fmt.Println("新节点补充")
			nodes = append(nodes, nodeNew)
		}

		//var nodeNow *node
		//nodeNow.val = nums[i]
		//nodeNow := &node{
		//	val:  nums[i],
		//	size: 1,
		//	next: nil,
		//}
		//nodes = append(nodes, nodeNow)
		fmt.Println("len nodes", len(nodes), max)
	}

	for _, c := range nodes {
		if c.size == max {
			for head := c; head != nil; {
				fmt.Print("val:", head.val)
				count[head.val]++
				head = head.next
			}
			fmt.Println()
		}

	}
	for i := 0; i < n; i++ {
		fmt.Println(nums[i], count[nums[i]])
	}

}
