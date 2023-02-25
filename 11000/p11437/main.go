package main

import (
	"bufio"
	"fmt"
	"os"
)

type tree struct {
	level  int
	value  int
	parent *tree
}

var arr = make(map[int]*tree)

func moveUP(cur *tree, level int) *tree {
	if cur.level == level {
		return cur
	}
	return moveUP(cur.parent, level)
}

func seachSameParent(a, b int) int {
	a1 := arr[a]
	b1 := arr[b]

	if a1.level < b1.level {
		b1 = moveUP(b1, a1.level)
	} else if a1.level > b1.level {
		a1 = moveUP(a1, b1.level)
	}

	// same level
	if a1.value == b1.value {
		return a1.value
	} else {
		for a1.value != b1.value {
			upLevel := a1.level - 1
			a1 = moveUP(a1, upLevel)
			b1 = moveUP(b1, upLevel)
		}
		return a1.value
	}
}

func main() {

	rd := bufio.NewReader(os.Stdin)
	wd := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(rd, &n)

	adj := make(map[int][]int)

	for i := 0; i < n-1; i++ {
		var a, b int
		fmt.Fscanln(rd, &a, &b)
		if vv, ex := adj[a]; ex {
			adj[a] = append(vv, b)
		} else {
			adj[a] = append(make([]int, 0), b)
		}
		if vv, ex := adj[b]; ex {
			adj[b] = append(vv, a)
		} else {
			adj[b] = append(make([]int, 0), a)
		}
	}

	// create tree
	q := make([]int, 1)
	q[0] = 1
	arr[1] = &tree{1, 1, nil}

	for len(q) > 0 {
		pp := q[0]
		q = q[1:]
		ptree := arr[pp]
		vv := adj[pp]
		for _, v := range vv {
			if _, ex2 := arr[v]; ex2 {
				continue
			}
			q = append(q, v)
			arr[v] = &tree{ptree.level + 1, v, ptree}
		}
	}

	var m int
	fmt.Fscanln(rd, &m)

	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscanln(rd, &a, &b)
		fmt.Fprintln(wd, seachSameParent(a, b))
	}

	wd.Flush()
}
