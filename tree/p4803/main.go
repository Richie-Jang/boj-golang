package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	testCase int
	dfs      func(before, cur int, adjs [][]int, visits []bool) bool
)

func main() {
	br := bufio.NewScanner(os.Stdin)
	testCase = 1

	getPair := func(s string) (int, int) {
		ss := strings.Split(s, " ")
		a, _ := strconv.Atoi(ss[0])
		b, _ := strconv.Atoi(ss[1])
		return a, b
	}

	createAdjs := func(nodes, count int) [][]int {
		adjs := make([][]int, nodes+1)
		for i := 1; i <= nodes; i++ {
			adjs[i] = make([]int, 0)
		}
		for i := 0; i < count; i++ {
			br.Scan()
			a, b := getPair(br.Text())
			adjs[a] = append(adjs[a], b)
			adjs[b] = append(adjs[b], a)
		}
		return adjs
	}

	// tree cycle checking
	dfs = func(before, cur int, adjs [][]int, visits []bool) bool {
		for i := 0; i < len(adjs[cur]); i++ {
			next := adjs[cur][i]
			if next == before {
				continue
			}
			if visits[next] {
				return false
			}
			visits[next] = true
			if !dfs(cur, next, adjs, visits) {
				return false
			}
		}
		return true
	}

	for br.Scan() {
		t := br.Text()
		a, b := getPair(t)
		if a == 0 && b == 0 {
			break
		}

		adjs := createAdjs(a, b)
		visits := make([]bool, a+1)
		result := 0
		for i := 1; i <= a; i++ {
			if !visits[i] {
				if dfs(0, i, adjs, visits) {
					result++
				}
			}
		}

		msg := ""
		switch {
		case result == 0:
			msg = "No trees."
		case result == 1:
			msg = "There is one tree."
		default:
			msg = fmt.Sprintf("A forest of %d trees.", result)
		}

		fmt.Printf("Case %d: %s\n", testCase, msg)
		testCase++
	}
}
