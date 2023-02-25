package main

import (
	"bufio"
	"fmt"
	"os"
)

var res = []string{}

func hanoi(n int, from, to, middle int) {
	if n == 1 {
		res = append(res, fmt.Sprintf("%d %d", from, to))
		return
	}

	hanoi(n-1, from, middle, to)
	res = append(res, fmt.Sprintf("%d %d", from, to))
	hanoi(n-1, middle, to, from)
}

func main() {
	var a int
	br := bufio.NewReader(os.Stdin)
	fmt.Fscanf(br, "%d\n", &a)

	hanoi(a, 1, 3, 2)

	fmt.Println(len(res))
	for _, a2 := range res {
		fmt.Println(a2)
	}
}
