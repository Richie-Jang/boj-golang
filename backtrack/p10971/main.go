package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	n      int
	adj    [][]int
	result int
)

func findMinRoute(start, cur, acc int, passed map[int]any) {

	// fmt.Println(start, cur, acc, passed)

	if len(passed) == n {
		finalWeight := adj[cur][start]
		if finalWeight == 0 {
			return
		}
		total := acc + finalWeight
		if total < result {
			result = total
		}
		return
	}

	// nexts
	for i, nx := range adj[cur] {
		if i == 0 {
			continue
		}
		if nx == 0 {
			continue
		}
		if _, ok := passed[i]; !ok {
			newWeight := acc + nx
			if newWeight > result {
				return
			}
			passed[i] = nil
			findMinRoute(start, i, newWeight, passed)
			delete(passed, i)
		}
	}
}

func main() {
	br := bufio.NewReader(os.Stdin)
	fmt.Fscanf(br, "%d\n", &n)
	adj = make([][]int, n+1)
	for i := 1; i <= n; i++ {
		adj[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			fmt.Fscanf(br, "%d", &adj[i][j])
		}
		_, _ = fmt.Fscanln(br)
	}

	result = math.MaxInt32

	for i := 1; i <= n; i++ {
		mm := make(map[int]any)
		mm[i] = nil
		findMinRoute(i, i, 0, mm)
		mm = nil
	}

	fmt.Println(result)

}
