package main

import (
	"fmt"
	"math"
)

var dp []int
var div, a, b int

func countNumber(v int) {
	loc := int(math.Log10(float64(v))) + 1
	for i := 1; i <= loc; i++ {
		div = int(math.Pow10(loc - i))
		a = v / div
		v = v % div
		dp[a] = dp[a] + 1
	}
}

func main() {
	var n = 1000000000
	dp = make([]int, 10)
	for i := 1; i <= n; i++ {
		countNumber(i)
	}
	fmt.Println(dp)
}
