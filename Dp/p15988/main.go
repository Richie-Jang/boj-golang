package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	MOD = 1_000_000_009
	MAX = 1_000_000
)

var (
	bw   *bufio.Writer
	t, n int
	arr  []int
)

func initArr() {
	arr = make([]int, MAX+1)
	arr[1] = 1
	arr[2] = 2
	arr[3] = 4
	for i := 4; i <= MAX; i++ {
		a1 := arr[i-1] % MOD
		a2 := arr[i-2] % MOD
		a3 := arr[i-3] % MOD
		arr[i] = (a1 + a2 + a3) % MOD
	}
}

func solve(br *bufio.Reader) {
	fmt.Fscanf(br, "%d\n", &n)
	bw.WriteString(fmt.Sprintf("%d\n", arr[n]))
}

func main() {
	br := bufio.NewReader(os.Stdin)
	bw = bufio.NewWriter(os.Stdout)
	fmt.Fscanf(br, "%d\n", &t)
	initArr()
	for i := 0; i < t; i++ {
		solve(br)
	}
	bw.Flush()
}
