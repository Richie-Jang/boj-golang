package main

import (
	"bufio"
	"fmt"
	"os"
)

var n, d, k, c int
var dishes []int
var table [3001]int

func getCount(cur int) int {
	res := 0
	rr := make(map[int]any)
	for i := cur; i < cur+k; i++ {
		if _, ok := rr[dishes[i]]; ok {
			continue
		}
		if table[dishes[i]] > 0 {
			res++
		}
		rr[dishes[i]] = nil
	}

	if _, ok := rr[c]; !ok {
		if table[c] > 0 {
			res++
		}
	}
	rr = nil
	//fmt.Println("Count", res)
	return res
}

func solve(curIdx, curCount int) int {
	endIdx := (curIdx + k - 1) % n
	table[dishes[curIdx-1]]--
	if table[dishes[curIdx-1]] == 0 {
		curCount--
	}
	table[dishes[endIdx]]++
	if table[dishes[endIdx]] == 1 {
		curCount++
	}

	//fmt.Println("curIdx", curIdx, "endIdx", endIdx, "Count", curCount)
	return curCount
}

func main() {
	br := bufio.NewReader(os.Stdin)
	fmt.Fscanf(br, "%d %d %d %d\n", &n, &d, &k, &c)
	dishes = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(br, "%d\n", &dishes[i])
	}
	// init
	for i := 0; i < k; i++ {
		table[dishes[i]]++
	}

	table[c]++

	var max, prev int = getCount(0), 0
	prev = max
	for i := 1; i < n; i++ {
		v := solve(i, prev)
		if max < v {
			max = v
		}
		prev = v
	}

	fmt.Println(max)

}
