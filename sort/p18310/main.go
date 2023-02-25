package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	n   int
	arr []int
)

func abs(v int) int {
	if v < 0 {
		return v * -1
	}
	return v
}

func computeDist(idx int) int {
	antPos := arr[idx]
	sum := 0
	for _, v := range arr {
		sum += abs(antPos - v)
	}
	return sum
}

type t2 struct {
	idx, value int
}

func solve() {
	midIdx := (n - 1) / 2
	ck1 := midIdx - 1
	ck2 := midIdx
	ck3 := midIdx + 1

	a1 := computeDist(ck1)
	a2 := computeDist(ck2)
	a3 := computeDist(ck3)

	nArr := []t2{{ck1, a1}, {ck2, a2}, {ck3, a3}}
	sort.Slice(nArr, func(i, j int) bool {
		a1 := nArr[i]
		a2 := nArr[j]
		return a1.value < a2.value
	})
	fmt.Println(arr[nArr[0].idx])
}

func main() {
	br := bufio.NewReader(os.Stdin)
	fmt.Fscanf(br, "%d\n", &n)

	arr = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscanf(br, "%d", &arr[i])
	}
	fmt.Fscanln(br)

	if n == 1 || n == 2 {
		fmt.Println(arr[0])
		return
	}

	sort.Ints(arr)
	solve()
}
