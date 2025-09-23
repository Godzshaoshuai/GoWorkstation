package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func compare(a, b int) {
	if a > b {
		fmt.Println(">")
	} else if a < b {
		fmt.Println("<")
	} else {
		fmt.Println("=")
	}

}
func power(a, b int) int {
	if a == 0 {
		return 0
	}
	if b == 0 {
		return 1
	}
	result := 1
	for i := 0; i < b; i++ {
		result *= a
	}
	return result
}

func main() {
	n := 0
	fmt.Scan(&n)
	in := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		in.Scan()
		str := in.Text()
		s := strings.Split(str, " ")
		a, _ := strconv.Atoi(s[0])
		b, _ := strconv.Atoi(s[1])
		c := s[2]
		if c == "/" {
			compare(a/b, b/a)
		} else if c == "*" {
			compare(a*b, b*a)
		} else if c == "+" {
			compare(a+b, b+a)
		} else if c == "-" {
			compare(a-b, b-a)
		} else if c == "^" {
			//fmt.Println(a, b, power(a, b), power(b, a))

			if float64(b)*math.Log(float64(a)) > float64(a)*math.Log(float64(b)) {
				fmt.Println(">")
			} else if float64(b)*math.Log(float64(a)) < float64(a)*math.Log(float64(b)) {
				fmt.Println("<")
			} else {
				fmt.Println("=")
			}
		}

	}

}
