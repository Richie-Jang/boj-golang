package main

import (
	"bufio"
	"fmt"
	"os"
)

var s string
var stack []int

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	s = sc.Text()
	// init
	stack = make([]int, 0)
	answer := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if i-last == 1 {
				answer += len(stack)
			} else {
				answer++
			}
		}
	}

	fmt.Println(answer)
}
