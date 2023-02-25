package main

import (
	"bufio"
	"fmt"
	"os"
)

// MAX : Array max
const MAX int = 200_002

var n int
var inArr []int

func main() {

	br := bufio.NewReader(os.Stdin)
	bw := bufio.NewWriter(os.Stdout)

	fmt.Fscan(br, &n)
	inArr = make([]int, n)
	maxValue := -1
	for i := 0; i < n; i++ {
		fmt.Fscan(br, &inArr[i])
		if maxValue < inArr[i] {
			maxValue = inArr[i]
		}
	}

	res := make([]int, MAX)
	for i := 0; i < MAX; i++ {
		res[i] = -1
	}

	ck := make(map[int]bool)
	ck[0] = true

	for i := 1; i < n; i++ {
		v := inArr[i]
		if _, ex := ck[v]; !ex {
			ck[v] = true
			prev := inArr[i-1]
			res[v] = prev
		}
	}

	fmt.Fprintln(bw, maxValue+1)
	for i := 0; i <= maxValue; i++ {
		fmt.Fprintf(bw, "%d ", res[i])
	}

	fmt.Fprintln(bw)
	bw.Flush()

}
