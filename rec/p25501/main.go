package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

var (
	n int
)

func rec(s string, left, right int, count int) (int, int) {
	if left >= right {
		return 1, count
	}
	if s[left] != s[right] {
		return 0, count
	}
	return rec(s, left+1, right-1, count+1)
}

func main() {

	br := bufio.NewReader(os.Stdin)
	fmt.Fscanf(br, "%d\n", &n)
	var s string
	buf := new(bytes.Buffer)
	for i := 0; i < n; i++ {
		fmt.Fscanf(br, "%s\n", &s)
		isPalindrome, count := rec(s, 0, len(s)-1, 1)
		buf.WriteString(fmt.Sprintf("%d %d\n", isPalindrome, count))
	}

	fmt.Print(buf.String())
	buf = nil
}
