package main

import (
	"bufio"
	"fmt"
	"os"
)

const MAX = 500_001

var (
	n    int
	adjs map[int][]int
)

type item struct {
	node, pos int
}

func solve() {
	q := []item{{1, 0}}
	visits := make([]bool, n+1)
	visits[1] = true

	pathCount := 0

	for len(q) > 0 {
		c := q[0]
		q = q[1:]

		ckcount := 0
		for _, p := range adjs[c.node] {
			if visits[p] {
				continue
			}
			visits[p] = true
			q = append(q, item{p, c.pos + 1})
			ckcount++
		}

		if ckcount == 0 {
			pathCount += c.pos
		}
	}

	if pathCount%2 != 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func main() {

	br := bufio.NewReader(os.Stdin)
	fmt.Fscanf(br, "%d\n", &n)
	var a, b int
	adjs = make(map[int][]int)
	for i := 0; i < n-1; i++ {
		fmt.Fscanf(br, "%d %d\n", &a, &b)
		if v, ok := adjs[a]; !ok {
			adjs[a] = []int{b}
		} else {
			adjs[a] = append(v, b)
		}
		if v, ok := adjs[b]; !ok {
			adjs[b] = []int{a}
		} else {
			adjs[b] = append(v, a)
		}
	}

	solve()

}
