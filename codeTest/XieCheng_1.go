package main

import (
	"fmt"
)

func fun1(v int) {
	if v < 160 || v > 350 {
		fmt.Println("other")
	} else if v >= 160 && v < 200 {
		fmt.Println("D")
	} else if v >= 200 && v < 250 {
		fmt.Println("D C")
	} else if v == 250 {
		fmt.Println("G D C")
	} else if v > 250 && v <= 300 {
		fmt.Println("G C")
	} else {
		fmt.Println("G")
	}
}

func main() {
	var n, v int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&v)
		fun1(v)
	}
}
