package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n, m int
	arr  [302]int
)

const SP = 100

func main() {
	br := bufio.NewReader(os.Stdin)
	fmt.Fscanf(br, "%d\n", &n)
	var k int
	for i := 0; i < n; i++ {
		fmt.Fscanf(br, "%d", &k)
		arr[k+SP]++
	}
	fmt.Fscanln(br)
	fmt.Fscanf(br, "%d\n", &m)

	fmt.Println(arr[m+SP])
}
