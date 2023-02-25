package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n int
	v int
	s string
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wd := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(rd, &n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(rd, &v, &s)
		a1 := s[:v-1]
		a2 := s[v:]
		fmt.Fprintln(wd, a1+a2)
	}
}
