package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	n, m   int
	visits []int
	adjs   map[int][]int
	ret    int
)

type pair struct {
	a, b int
}

func dfs(cur, count int) {
	if count == 5 {
		ret = 1
		return
	}
	visits[cur] = 1
	//fmt.Println("Cur", cur, "Count", count, visits)
	nps, ex := adjs[cur]
	if ex {
		for _, np := range nps {
			if visits[np] == 0 {
				//fmt.Print("Going => ", np, " ")
				dfs(np, count+1)
				if ret == 1 {
					return
				}
			}
		}
	}
	visits[cur] = 0
}

func main() {
	// 	test1 := `8 8
	// 1 7
	// 3 7
	// 4 7
	// 3 4
	// 4 6
	// 3 5
	// 0 4
	// 2 7`

	rd := bufio.NewReader(os.Stdin)
	//	rd := bufio.NewReader(strings.NewReader(test1))
	fmt.Fscanln(rd, &n, &m)
	visits = make([]int, n)
	adjs = make(map[int][]int)
	var a, b int
	for i := 0; i < m; i++ {
		fmt.Fscanln(rd, &a, &b)
		if v, ex := adjs[a]; ex {
			adjs[a] = append(v, b)
		} else {
			adjs[a] = []int{b}
		}
		if v, ex := adjs[b]; ex {
			adjs[b] = append(v, a)
		} else {
			adjs[b] = []int{a}
		}
	}

	// start 0
	isDone := false

	nodes := []pair{}
	for i := 0; i < n; i++ {
		if v, ex := adjs[i]; ex {
			nodes = append(nodes, pair{i, len(v)})
		}
	}
	sort.SliceStable(nodes, func(c, d int) bool {
		a1 := nodes[c]
		a2 := nodes[d]
		return a1.b < a2.b
	})
	// start node check
	for _, r := range nodes {
		dfs(r.a, 1)
		if ret == 1 {
			isDone = true
			break
		}
	}

	if !isDone {
		fmt.Println("0")
	} else {
		fmt.Println("1")
	}

}
