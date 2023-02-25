package main

import (
	"bufio"
	"fmt"
	"os"
)

var t, n, m int
var adj [][]int
var groups []int

func bfs(start int) bool {
	queue := make([]int, 1)
	queue[0] = start
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		state := groups[v]
		for j := 0; j < len(adj[v]); j++ {
			next := adj[v][j]
			if groups[next] == 0 {
				groups[next] = -1 * state
				queue = append(queue, next)
			} else if groups[next] == state {
				return false
			}
		}
	}
	return true
}

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Fscanf(r, "%d\n", &t)

	for i := 1; i <= t; i++ {
		fmt.Fscanf(r, "%d %d\n", &n, &m)
		adj = make([][]int, n+1)
		groups = make([]int, n+1)
		for j := 0; j <= n; j++ {
			adj[j] = make([]int, 0)
			groups[j] = 0
		}
		var a, b int
		for j := 1; j <= m; j++ {
			fmt.Fscanf(r, "%d %d\n", &a, &b)
			adj[a] = append(adj[a], b)
			adj[b] = append(adj[b], a)
		}

		var result = true
		for j := 1; j <= n; j++ {
			if groups[j] == 0 {
				groups[j] = 1
				if !bfs(j) {
					result = false
					break
				}
			}
		}

		if result {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
