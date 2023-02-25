package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var n, m int
var depths [32001]int
var data [32001][]int

func solve_problem() {
	q := make([]int, 0)
	for i := 1; i <= n; i++ {
		if depths[i] == 0 {
			q = append(q, i)
		}
	}

	var res = make([]string, 0)

	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		res = append(res, strconv.Itoa(node))
		for j := 0; j < len(data[node]); j++ {
			next := data[node][j]
			depths[next]--
			if depths[next] == 0 {
				q = append(q, next)
			}
		}
	}

	fmt.Println(strings.Join(res, " "))
}

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Fscanf(r, "%d %d\n", &n, &m)
	var a, b int
	for i := 1; i <= n; i++ {
		data[i] = make([]int, 0)
		depths[i] = 0
	}
	for i := 1; i <= m; i++ {
		fmt.Fscanf(r, "%d %d\n", &a, &b)
		data[a] = append(data[a], b)
		depths[b]++
	}

	solve_problem()
}
