package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var larr, harr []int

func main() {
	// 	var input = `14 5
	// 1
	// 3
	// 4
	// 2
	// 2
	// 4
	// 3
	// 4
	// 3
	// 3
	// 3
	// 2
	// 3
	// 3`

	rd := bufio.NewReader(os.Stdin)
	var n, h int
	fmt.Fscanln(rd, &n, &h)
	var half = n / 2
	larr = make([]int, half)
	harr = make([]int, half)

	for i := 0; i < half; i++ {
		fmt.Fscan(rd, &larr[i])
		fmt.Fscan(rd, &harr[i])
	}

	sort.Ints(larr)
	sort.Ints(harr)

	var cnt = 0
	var space = n * 2

	for i := 1; i <= h; i++ {
		ls := sort.SearchInts(larr, i)
		hs := sort.SearchInts(harr, h-i+1)
		t := half - ls + half - hs
		if t < space {
			space = t
			cnt = 1
		} else if t == space {
			cnt++
		}
	}

	fmt.Println(space, cnt)
}
