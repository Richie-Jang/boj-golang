package main

import "fmt"

var n int

func main() {
	fmt.Scanf("%d\n", &n)
	dp := make([]int, 10_000_001)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3
	dp[4] = 5
	dp[5] = 8
	dp[6] = 13

	for i := 7; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % 10
	}

	fmt.Println(dp[n] % 10)

}
