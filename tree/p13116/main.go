package main

import (
	"bufio"
	"fmt"
	"os"
)

var n int

func getAnswer(a, b int) int {
	if a == b {
		return a
	}
	if a < b {
		return getAnswer(a, b/2)
	} else {
		return getAnswer(a/2, b)
	}
}

func main() {
	br := bufio.NewReader(os.Stdin)
	bw := bufio.NewWriter(os.Stdout)
	fmt.Fscan(br, &n)

	a, b := 0, 0
	for i := 0; i < n; i++ {
		fmt.Fscan(br, &a, &b)
		ans := getAnswer(a, b)
		fmt.Fprintln(bw, ans*10)
	}
	bw.Flush()
}
