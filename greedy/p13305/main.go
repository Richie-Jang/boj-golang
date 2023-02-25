package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n      int
	dists  []int
	prices []int
)

func getDistSum(start, count, acc int) int {
	if count == 0 {
		return acc
	}
	return getDistSum(start+1, count-1, acc+dists[start])
}

func loop(cur, acc int) int {
	if cur == n-1 {
		return acc
	}
	curPrice := prices[cur]
	nextPos := cur + 1
	for nextPos < n-1 && (curPrice <= prices[nextPos]) {
		nextPos++
	}
	dist := getDistSum(cur, nextPos-cur, 0)
	return loop(nextPos, acc+dist*curPrice)
}

func main() {
	br := bufio.NewReader(os.Stdin)
	fmt.Fscanf(br, "%d\n", &n)
	dists = make([]int, n-1)
	prices = make([]int, n)
	for i := 0; i < n-1; i++ {
		fmt.Fscanf(br, "%d", &dists[i])
	}
	fmt.Fscanln(br)
	for i := 0; i < n; i++ {
		fmt.Fscanf(br, "%d", &prices[i])
	}
	fmt.Fscanln(br)

	fmt.Println(loop(0, 0))
}
