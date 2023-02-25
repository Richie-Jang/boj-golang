package main

import (
	"bufio"
	"fmt"
	"os"
)

func makePi(s string) []int {
	l := len(s)
	j := 0
	res := make([]int, l)
	for i := 1; i < l; i++ {
		for j > 0 && s[i] != s[j] {
			j = res[j-1]
		}
		if s[i] == s[j] {
			j++
			res[i] = j
		}
	}
	return res
}

func main() {
	br := bufio.NewReader(os.Stdin)
	var s, l string
	fmt.Fscanln(br, &s)
	fmt.Fscanln(br, &l)
	table := makePi(s)

	fmt.Println("table: ", table)
	var j, ll int
	ll = len(l)
	ans := 0
	// kmp
	for i := 0; i < len(s); i++ {
		for j > 0 && s[i] != l[j] {
			prevj := j
			j = table[j-1]
			fmt.Println("NotMatch i J :", i, prevj, "->", j)
		}
		if s[i] == l[j] {
			if j == ll-1 {
				ans = 1
				break
			} else {
				j++
			}
		}
	}

	fmt.Println(ans)
}
