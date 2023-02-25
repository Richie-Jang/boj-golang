package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const MINVAL = -1_000_000_002

type line struct {
	st, ed int
}

var (
	n     int
	lines []line = make([]line, 0, 1_000_001)
)

func main() {
	br := bufio.NewReader(os.Stdin)
	fmt.Fscanf(br, "%d\n", &n)
	var a, b int

	for i := 0; i < n; i++ {
		fmt.Fscanf(br, "%d %d\n", &a, &b)
		lines = append(lines, line{a, b})
	}

	sort.Slice(lines, func(i, j int) bool {
		return lines[i].st < lines[j].st
	})

	var res int
	var Left, Right = MINVAL, MINVAL

	for _, l := range lines {
		if l.st <= Right {
			mm := Right
			if l.ed > mm {
				mm = l.ed
			}
			Right = mm
		} else {
			res += Right - Left
			Left = l.st
			Right = l.ed
		}
	}
	res += Right - Left

	fmt.Println(res)
}
