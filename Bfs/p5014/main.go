package main

import (
	"fmt"
)

var (
	f, s, g, u, d int
	dp            [1000001]int
)

func solve() {
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}
	q := make([]int, 1)
	q[0] = s
	dp[s] = 0
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		if p == g {
			break
		}
		// up
		up := p + u
		down := p - d
		ds := []int{up, down}
		for i := 0; i < 2; i++ {
			if ds[i] < 1 || ds[i] > 1000000 || dp[ds[i]] != -1 {
				continue
			}
			q = append(q, ds[i])
			dp[ds[i]] = dp[p] + 1
		}
	}
	if dp[g] == -1 {
		fmt.Println("use the stairs")
	} else {
		fmt.Println(dp[g])
	}
}

func main() {
	fmt.Scanf("%d %d %d %d %d", &f, &s, &g, &u, &d)
	solve()
}
