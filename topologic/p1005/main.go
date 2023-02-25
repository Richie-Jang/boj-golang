package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	t int
)

func solve(br *bufio.Reader) {
	var n, k int
	fmt.Fscanf(br, "%d %d\n", &n, &k)

	delays := make([]int, n+1)
	adjs := make(map[int][]int)
	depths := make(map[int]int)
	costs := make(map[int]int)
	for i := 1; i <= n; i++ {
		fmt.Fscanf(br, "%d", &delays[i])
	}
	fmt.Fscanln(br)
	var a, b, goal int

	for i := 0; i < k; i++ {
		fmt.Fscanf(br, "%d %d\n", &a, &b)
		adjs[a] = append(adjs[a], b)
		depths[b] += 1
	}
	fmt.Fscanf(br, "%d\n", &goal)

	q := make([]int, 0)
	for i := 1; i <= n; i++ {
		if depths[i] == 0 {
			q = append(q, i)
		}
	}

	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		// compute delays
		costs[v] = costs[v] + delays[v]
		nvs := adjs[v]
		for _, nv := range nvs {
			costs[nv] = max(costs[v], costs[nv])
			depths[nv]--
			if depths[nv] == 0 {
				q = append(q, nv)
			}
		}
		if v == goal {
			break
		}
	}
	fmt.Println(costs[goal])

}

func main() {
	br := bufio.NewReader(os.Stdin)
	fmt.Fscanf(br, "%d\n", &t)
	for i := 0; i < t; i++ {
		solve(br)
	}
}
