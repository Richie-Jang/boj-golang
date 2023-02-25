package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	MAX = 51
)

var (
	n   int
	dd  [MAX]int
	adj [MAX][]int
	pow [MAX][3]int
)

func dfs(cur int, vs []bool) bool {
	if vs[cur] {
		return false
	}
	vs[cur] = true
	for i := 0; i < len(adj[cur]); i++ {
		v := adj[cur][i]
		//fmt.Println("cur", cur, "v", v, "adj", adj[cur])
		if dd[v] == 0 || dfs(dd[v], vs) {
			dd[v] = cur
			//fmt.Println("call:", dd[v], cur)
			return true
		}
	}
	return false
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	fmt.Fscanln(rd, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscanln(rd, &pow[i][0], &pow[i][1], &pow[i][2])
		adj[i] = []int{}
	}

	for i := 1; i <= n-1; i++ {
		for j := i + 1; j <= n; j++ {
			if pow[i][0] >= pow[j][0] && pow[i][1] >= pow[j][1] && pow[i][2] >= pow[j][2] {
				adj[i] = append(adj[i], j)
			} else if pow[i][0] <= pow[j][0] && pow[i][1] <= pow[j][1] && pow[i][2] <= pow[j][2] {
				adj[j] = append(adj[j], i)
			}
		}
	}

	//fmt.Println("adj:", adj[1:n+1])
	//fmt.Println("dd:", dd[1:n+1])

	var res = 0
	for k := 1; k <= 2; k++ {
		for i := 1; i <= n; i++ {
			var vs = make([]bool, n+1)
			for j := 1; j <= n; j++ {
				vs[i] = false
			}
			if dfs(i, vs) {
				res++
			}
		}
	}
	fmt.Println(n - res)
}
