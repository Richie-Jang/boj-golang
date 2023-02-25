package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	MAX = 40001
)

type node struct {
	value  int
	parent *node
	level  int
	dist   int
}

type pair struct {
	tonode int
	dist   int
}

var arr = make(map[int]*node)

func solve(a, b int) int {
	a1 := arr[a]
	a2 := arr[b]
	res := 0
	for a1.level > a2.level {
		res += a1.dist
		a1 = a1.parent
	}
	for a1.level < a2.level {
		res += a2.dist
		a2 = a2.parent
	}
	for a1.value != a2.value {
		res += (a1.dist + a2.dist)
		a1 = a1.parent
		a2 = a2.parent
	}
	return res
}

func main() {

	rd := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(rd, &n)

	adj := make(map[int][]pair)

	for i := 0; i < n-1; i++ {
		var a, b, c int
		fmt.Fscanln(rd, &a, &b, &c)
		if vv, ex := adj[a]; ex {
			adj[a] = append(vv, pair{b, c})
		} else {
			adj[a] = []pair{{b, c}}
		}

		if vv, ex := adj[b]; ex {
			adj[b] = append(vv, pair{a, c})
		} else {
			adj[b] = []pair{{a, c}}
		}
	}

	// setup
	q := []int{1}
	arr[1] = &node{1, nil, 1, 0}
	for len(q) > 0 {
		p := q[0]
		curNode := arr[p]
		q = q[1:]
		for _, np := range adj[p] {
			if _, ex := arr[np.tonode]; ex {
				continue
			}
			q = append(q, np.tonode)
			arr[np.tonode] = &node{np.tonode, curNode, curNode.level + 1, np.dist}
		}
	}

	var m int
	fmt.Fscanln(rd, &m)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscanln(rd, &a, &b)
		fmt.Println(solve(a, b))
	}

}
