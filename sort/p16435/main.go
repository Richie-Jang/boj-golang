package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	n, k    int
	heights []int
)

func main() {
	br := bufio.NewReader(os.Stdin)
	fmt.Fscanf(br, "%d %d\n", &n, &k)
	heights = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(br, "%d", &heights[i])
	}
	fmt.Fscanln(br)
	sort.Ints(heights)

	for i := 0; i < n; i++ {
		h := heights[i]
		if h <= k {
			k++
			continue
		}
		break
	}

	fmt.Println(k)
}
