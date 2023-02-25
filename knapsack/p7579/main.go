package main

import (
	"bufio"
	"fmt"
	"os"
)

const MAX int = 101

var (
	n, m int
	dp   [MAX][MAX * MAX]int
	app  [101]int
	cost [101]int
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

	br := bufio.NewReader(os.Stdin)
	fmt.Fscan(br, &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Fscan(br, &app[i])
	}
	memSum := 0
	for i := 1; i <= n; i++ {
		fmt.Fscan(br, &cost[i])
		memSum += cost[i]
	}

	for i := 1; i <= n; i++ {
		curApp := app[i]
		for j := 1; j <= memSum; j++ {
			dp[i][j] = dp[i-1][j]
			if j-cost[i] >= 0 {
				dp[i][j] = max(dp[i][j], dp[i-1][j-cost[i]]+curApp)
			}
		}
	}

	// print
	// for i := 1; i <= n; i++ {
	// 	for j := 1; j <= memSum; j++ {
	// 		fmt.Print(dp[i][j], " ")
	// 	}
	// 	fmt.Println("")
	// }

	for i := 1; ; i++ {
		if dp[n][i] >= m {
			fmt.Println(i)
			break
		}
	}

}
