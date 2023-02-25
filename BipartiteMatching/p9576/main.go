package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	t, n, m int
	p       [2002]int
	adj     [2002][]int
)

func dfs(cur int, vs []bool) bool {
	if vs[cur] {
		return false
	}
	vs[cur] = true
	for i := 0; i < len(adj[cur]); i++ {
		v := adj[cur][i]
		if p[v] == 0 || dfs(p[v], vs) {
			p[v] = cur
			return true
		}
	}
	return false
}

func solve() {
	res := 0
	var vs = make([]bool, 2002)
	for i := 1; i <= n; i++ {
		for k := 1; k <= 2001; k++ {
			vs[k] = false
		}
		if dfs(i, vs) {
			res++
		}
	}
	fmt.Println(res)
}

func readingData(rd *bufio.Reader) {
	// clear
	for i := 0; i <= 2001; i++ {
		adj[i] = []int{}
		p[i] = 0
	}

	fmt.Fscanln(rd, &n, &m)
	var a, b int
	for i := 1; i <= m; i++ {
		fmt.Fscanln(rd, &a, &b)
		for j := a; j <= b; j++ {
			adj[j] = append(adj[j], i+n)
		}
	}
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	fmt.Fscanln(rd, &t)
	for i := 1; i <= t; i++ {
		readingData(rd)
		//fmt.Println(adj[1 : n+m+1])
		solve()
	}
}
