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

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	br := bufio.NewReader(os.Stdin)
	fmt.Fscanf(br, "%d\n", &n)
	var a, b int
	buf := new(bytes.Buffer)
	for i := 0; i < n; i++ {
		fmt.Fscanf(br, "%d %d\n", &a, &b)
		buf.WriteString(fmt.Sprintf("%d\n", gcd(a, b)))
	}

	fmt.Print(buf.String())
}
