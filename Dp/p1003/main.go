package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 41

var t, n int

type Pair struct {
	_1 int
	_2 int
}

func solve(n int) Pair {
	dp := [N]Pair{
		Pair{1, 0},
		Pair{0, 1},
		Pair{1, 1},
	}
	if n <= 2 {
		return dp[n]
	}

	if dp[n] != (Pair{0, 0}) {
		return dp[n]
	}

	for i := 3; i <= n; i++ {
		a1 := dp[i-1]._1
		b1 := dp[i-1]._2
		a2 := dp[i-2]._1
		b2 := dp[i-2]._2
		dp[i] = Pair{a1 + a2, b1 + b2}
	}
	return dp[n]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &t)
	for i := 1; i <= t; i++ {
		fmt.Fscanln(reader, &n)
		res := solve(n)
		fmt.Println(res._1, res._2)
	}
}
