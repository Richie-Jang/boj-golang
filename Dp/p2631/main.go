package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(arr []int) {
	n := len(arr)
	var dp = make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			left_value := arr[j]
			right_value := arr[i]
			if left_value < right_value {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
	}

	var max_value = 0
	for i := 0; i < n; i++ {
		if dp[i] > max_value {
			max_value = dp[i]
		}
	}

	fmt.Println(n - max_value)
}

func main() {
	var n, d int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &n)
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d\n", &d)
		arr[i] = d
	}

	solve(arr)
}
