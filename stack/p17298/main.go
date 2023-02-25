package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n   int
	arr []int
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	fmt.Fscanln(rd, &n)
	arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &arr[i])
	}

	stack := make([]int, 0)
	data := make([]int, n)

	for i := 0; i < n; i++ {
		for len(stack) > 0 {
			head := stack[len(stack)-1]
			if arr[head] < arr[i] {
				data[head] = arr[i]
				stack = stack[0 : len(stack)-1]
			} else {
				break
			}
		}
		stack = append(stack, i)
	}

	for i := 0; i < n; i++ {
		v := data[i]
		if v == 0 {
			v = -1
		}
		if i == n-1 {
			fmt.Println(v)
		} else {
			fmt.Print(v, " ")
		}

	}

}
