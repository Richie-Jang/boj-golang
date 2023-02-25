package main

import (
	"fmt"
)

var dp [201][201]int

func loop(n, k int) int {
	if k == 1 {
		return 1
	}
	var g = dp[n][k]
	if g > 0 {
		return g
	}
	for i := 0; i <= n; i++ {
		g = (g + loop(n-i, k-1)) % 1000000000
		dp[n][k] = g
	}
	return dp[n][k]
}

func main() {
	var n, k int
	fmt.Scanf("%d %d\n", &n, &k)
	fmt.Println(loop(n, k))
}
