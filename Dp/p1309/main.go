package main

import (
	"fmt"
)

var dp [100001][3]int

const DIV = 9901

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	dp[1][0] = 1
	dp[1][1] = 1
	dp[1][2] = 1

	for i := 2; i <= n; i++ {
		dp[i][0] = (dp[i-1][0] + dp[i-1][1] + dp[i-1][2]) % DIV
		dp[i][1] = (dp[i-1][0] + dp[i-1][2]) % DIV
		dp[i][2] = (dp[i-1][0] + dp[i-1][1]) % DIV
	}

	res := (dp[n][0] + dp[n][1] + dp[n][2]) % DIV
	fmt.Println(res)

}
