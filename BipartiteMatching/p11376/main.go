package main

import (
	"bufio"
	"fmt"
	"os"
)

type fc struct {
	f, c int
}

type pair struct {
	_1, _2 int
}

var (
	n, m  int
	adj   map[int][]int
	fcmap map[pair]fc
)

func main() {
	adj = make(map[int][]int)
	fcmap = make(map[pair]fc)
	rd := bufio.NewReader(os.Stdin)
	fmt.Fscanf(rd, "%d %d\n", &n, &m)
	var tmp string
	for i := 1; i <= n; i++ {
		adj[i] = make([]int, 0)
		var wc, cc int
		fmt.Fscan(rd, &wc)
		for j := 0; j < wc; j++ {
			fmt.Fscan(rd, &cc)
			cc += n
			adj[i] = append(adj[i], cc)
			adj[cc] = append(adj[cc], i)
		}
		fmt.Fscanln(rd, tmp)
	}

	fmt.Println(adj)
}
