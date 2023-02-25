package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n, m         int
	visit        [201]bool
	cowFromHouse [201]int
	adj          [201][]int
	dfs          func(int) bool
)

func solve() int {
	dfs = func(cur int) bool {
		if visit[cur] {
			return false
		}
		visit[cur] = true
		for j := 0; j < len(adj[cur]); j++ {
			next := adj[cur][j]
			if cowFromHouse[next] == 0 || dfs(cowFromHouse[next]) {
				cowFromHouse[next] = cur
				return true
			}
		}
		return false
	}

	var res = 0

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			visit[j] = false
		}
		if dfs(i) {
			res++
		}
	}
	return res
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	fmt.Fscanf(rd, "%d %d\n", &n, &m)
	var c int
	var tmp string
	for i := 1; i <= n; i++ {
		fmt.Fscan(rd, &c)
		adj[i] = make([]int, c)
		for j := 0; j < c; j++ {
			fmt.Fscan(rd, &adj[i][j])
		}
		fmt.Fscanln(rd, &tmp)
	}

	fmt.Println(solve())
}
