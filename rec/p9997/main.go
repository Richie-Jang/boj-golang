package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n          int
	bitarr     []int
	res        int
	allchecked int = (1 << 26) - 1
)

func solve(idx, sum int) {
	if idx == n {
		if sum == allchecked {
			res++
		}
		return
	}

	solve(idx+1, sum)
	solve(idx+1, sum|bitarr[idx])
}

func main() {
	br := bufio.NewReader(os.Stdin)
	fmt.Fscanf(br, "%d\n", &n)
	bitarr = make([]int, 0, n)
	var s string
	for i := 0; i < n; i++ {
		fmt.Fscanf(br, "%s\n", &s)
		var sum int
		for j := 0; j < len(s); j++ {
			sum |= (1 << (s[j] - 'a'))
		}
		bitarr = append(bitarr, sum)
	}

	solve(0, 0)
	fmt.Println(res)
}
