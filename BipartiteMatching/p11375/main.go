package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n, m          int
	emp           [1001][]int
	visit         [1001]bool
	personFromPos [1001]int
)

func countBipartiteMatch() int {
	var dfs func(int) bool

	dfs = func(cur int) bool {
		if visit[cur] {
			return false
		}
		visit[cur] = true
		for k := 0; k < len(emp[cur]); k++ {
			nextPos := emp[cur][k]
			if personFromPos[nextPos] == 0 || dfs(personFromPos[nextPos]) {
				personFromPos[nextPos] = cur
				return true
			}
		}
		return false
	}

	var res = 0
	for i := 1; i <= n; i++ {
		// init visit
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
	var tmp string
	for i := 1; i <= n; i++ {
		var wc int
		fmt.Fscan(rd, &wc)
		emp[i] = make([]int, wc)
		for j := 0; j < wc; j++ {
			fmt.Fscan(rd, &emp[i][j])
		}
		fmt.Fscanln(rd, tmp)
	}
	fmt.Println(countBipartiteMatch())
}
