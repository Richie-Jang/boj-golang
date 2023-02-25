package main

import (
	"fmt"
	"sort"
	"strings"
)

type Pair struct {
	st, ed int
}

var (
	s     string
	pairs []Pair
	rec   func(idx, count int)
)

func solve() {
	res := make(map[string]any)
	visited := make([]bool, len(s))
	bs := make([]bool, len(s))
	rec = func(idx, count int) {
		if count >= 1 {
			nstr := []byte(s)
			for i := 0; i < len(s); i++ {
				if bs[i] {
					nstr[i] = '^'
				}
			}
			newstr := strings.Replace(string(nstr), "^", "", -1)
			res[newstr] = nil
		}

		for i := idx; i < len(pairs); i++ {
			if visited[i] {
				continue
			}
			visited[i] = true
			p := pairs[i]
			bs[p.st] = true
			bs[p.ed] = true
			rec(idx+1, count+1)
			visited[i] = false
			bs[p.st] = false
			bs[p.ed] = false
		}
	}

	rec(0, 0)

	resbuf := make([]string, 0, len(res))
	for k := range res {
		resbuf = append(resbuf, k)
	}
	sort.Strings(resbuf)
	fmt.Println(strings.Join(resbuf, "\n"))
}

func main() {
	fmt.Scanln(&s)
	stack := make([]int, 0)
	pairs = make([]Pair, 0)
	for i, v := range s {
		if v == '(' {
			stack = append(stack, i)
			continue
		}
		if v == ')' {
			p := Pair{ed: i}
			p.st = stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			pairs = append(pairs, p)
		}
	}

	solve()
}
