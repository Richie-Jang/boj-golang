package main

import (
	"bufio"
	"fmt"
	"os"
)

var n int
var adj [501][]int
var depths [501]int
var times [501]int
var res [501]int

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func solve_problem() {

	q := make([]int, 0)

	for i := 1; i <= n; i++ {
		if depths[i] == 0 {
			res[i] = times[i]
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		for i := 0; i < len(adj[v]); i++ {
			next := adj[v][i]
			depths[next]--
			res[next] = max(res[next], res[v]+times[next])
			if depths[next] == 0 {
				q = append(q, next)
			}
		}

	}

	for i := 1; i <= n; i++ {
		fmt.Println(res[i])
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Fscanf(r, "%d\n", &n)

	var c, tmp int
	for i := 1; i <= n; i++ {
		adj[i] = make([]int, 0)
	}
	for i := 1; i <= n; i++ {
		res[i] = 0
		depths[i] = 0
		fmt.Fscanf(r, "%d", &times[i])
		for {
			fmt.Fscanf(r, "%d", &c)
			if c == -1 {
				break
			}
			adj[c] = append(adj[c], i)
			depths[i]++
		}
		fmt.Fscanln(r, &tmp)
	}

	solve_problem()

}
