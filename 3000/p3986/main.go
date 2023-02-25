package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n    int
	line string
)

func checkRun() bool {
	stack := make([]rune, 0)
	for _, v := range line {
		//fmt.Println("stack:", stack, "Cur:", v)
		count := len(stack)
		if count == 0 {
			stack = append(stack, v)
		} else {
			v2 := stack[count-1]
			if v == v2 {
				stack = stack[0 : count-1]
			} else {
				stack = append(stack, v)
			}
		}
	}

	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	br := bufio.NewReader(os.Stdin)
	fmt.Fscanf(br, "%d\n", &n)
	totalCount := 0
	for i := 0; i < n; i++ {
		fmt.Fscanf(br, "%s\n", &line)
		if checkRun() {
			totalCount += 1
		}
	}

	fmt.Println(totalCount)
}
