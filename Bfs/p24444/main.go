package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

const (
	MAX = 100_001
)

var (
	n, m, r int
	adjs    map[int][]int
)

func bfs() {
	var visits [MAX]bool
	visits[r] = true
	counterMap := make(map[int]int)
	// sort
	for _, v := range adjs {
		sort.Ints(v)
	}

	q := []int{r}
	count := 0
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		count++
		counterMap[v] = count
		for _, b := range adjs[v] {
			if !visits[b] {
				visits[b] = true
				q = append(q, b)
			}
		}
	}

	buf := new(bytes.Buffer)

	for i := 1; i <= n; i++ {
		if c, ok := counterMap[i]; !ok {
			buf.WriteString("0\n")
		} else {
			buf.WriteString(fmt.Sprintf("%d\n", c))
		}
	}

	fmt.Print(buf.String())
}

func main() {
	br := bufio.NewReader(os.Stdin)
	fmt.Fscanf(br, "%d %d %d\n", &n, &m, &r)
	var a, b int
	adjs = make(map[int][]int)
	for i := 0; i < m; i++ {
		fmt.Fscanf(br, "%d %d\n", &a, &b)
		if _, ok := adjs[a]; !ok {
			adjs[a] = []int{b}
		} else {
			adjs[a] = append(adjs[a], b)
		}
		if _, ok := adjs[b]; !ok {
			adjs[b] = []int{a}
		} else {
			adjs[b] = append(adjs[b], a)
		}
	}
	bfs()
}
