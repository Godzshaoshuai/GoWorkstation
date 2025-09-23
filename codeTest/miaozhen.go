package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func jishu(h1, m1, s1, h2, m2, s2 int) float64 {
	var hc float64 = 3600.0
	var mc float64 = 60.0
	//fmt.Println(h1, m1, s1, h2, m2, s2)
	totals1 := float64(h1)*hc + float64(m1)*mc + float64(s1)
	totals2 := float64(h2)*hc + float64(m2)*mc + float64(s2)
	//fmt.Println(totals1)
	//fmt.Println(totals2)
	if totals1 > totals2 {
		return float64(24*3600-totals1+totals2) / 60
	}
	return float64(totals2-totals1) / 60

}

func main() {
	n := 0
	fmt.Scan(&n)
	in := bufio.NewScanner(os.Stdin)
	var s []string
	in.Scan()
	str := in.Text()
	s = strings.Fields(str)

	for i := 0; i < n-1; i++ {
		stime1 := strings.Split(s[i], ":")
		stime2 := strings.Split(s[i+1], ":")
		h1, _ := strconv.Atoi(stime1[0])
		m1, _ := strconv.Atoi(stime1[1])
		s1, _ := strconv.Atoi(stime1[2])
		h2, _ := strconv.Atoi(stime2[0])
		m2, _ := strconv.Atoi(stime2[1])
		s2, _ := strconv.Atoi(stime2[2])
		res := strconv.FormatFloat(jishu(h1, m1, s1, h2, m2, s2), 'f', 10, 64)
		fmt.Print(res, " ")
	}

}
