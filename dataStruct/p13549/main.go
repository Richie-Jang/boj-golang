package main

import (
	"fmt"
)

const MAX = 100_000
const MAXLEN = 100_100

func solve(n, m int) {
	var vs [MAXLEN]int
	q := make([]int, 0)
	q = append(q, n)
	vs[n] = 1

	for len(q) > 0 {
		now := q[0]
		q = q[1:]
		if now == m {
			fmt.Println(vs[now] - 1)
			break
		}
		nd := now * 2
		if nd > 0 && nd <= MAX && vs[nd] == 0 {
			q = append(q, nd)
			vs[nd] = vs[now]
		}
		nd = now - 1
		if nd >= 0 && nd <= MAX && vs[nd] == 0 {
			q = append(q, nd)
			vs[nd] = vs[now] + 1
		}
		nd = now + 1
		if nd >= 0 && nd <= MAX && vs[nd] == 0 {
			q = append(q, nd)
			vs[nd] = vs[now] + 1
		}
	}
}

func main() {
	var n, m int
	fmt.Scanln(&n, &m)
	solve(n, m)
}
