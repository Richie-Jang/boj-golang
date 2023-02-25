package main

import (
	"fmt"
)

const MAX = 100_001

var dist, path [MAX]int

func main() {

	var n, m int
	fmt.Scanln(&n, &m)

	for i := range dist {
		dist[i] = -1
	}

	q := make([]int, 0)
	q = append(q, n)
	dist[n] = 0

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if cur == m {
			fmt.Println(dist[cur])
			ps := make([]int, 0)
			for cur != n {
				ps = append(ps, cur)
				cur = path[cur]
			}
			ps = append(ps, n)

			for i := len(ps) - 1; i >= 0; i-- {
				fmt.Printf("%d ", ps[i])
			}
			fmt.Println()
			break
		}

		for _, v := range []int{cur * 2, cur + 1, cur - 1} {
			if v < 0 || v >= MAX {
				continue
			}
			if dist[v] != -1 {
				continue
			}
			q = append(q, v)
			dist[v] = dist[cur] + 1
			path[v] = cur
		}
	}
}
