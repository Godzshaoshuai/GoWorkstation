package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	vmap := make(map[string]bool)
	//fmt.Println(s)
	startInde := 0
	for i := 0; i < len(s); i++ {
		if startInde != i && s[i] == 'd' && s[i+1] == 'p' {
			s1 := s[startInde:i]
			startInde = i + 2
			i++
			if !vmap[s1] {
				fmt.Println(s1)
				vmap[s1] = true
			}
		}
	}
	fmt.Print(len(vmap))

}
