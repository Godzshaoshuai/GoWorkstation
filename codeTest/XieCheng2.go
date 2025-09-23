package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, value int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	max := -1
	result := 0
	for i := 0; i < n; i++ {
		scanner.Scan()
		value, _ = strconv.Atoi(scanner.Text())
		if max < value {
			max = value
		}
		result = max - i - 1
		fmt.Fprint(writer, result)
		if i < n-1 {
			fmt.Fprint(writer, " ")
		}
	}
	fmt.Fprint(writer)

}
