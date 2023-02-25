package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

var (
	t, n int
	arr  []int
)

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func solve(buf *bytes.Buffer) {
	sort.Ints(arr)
	newArr := make([]int, n)
	front, rear := 0, n-1
	for i, a := range arr {
		if i%2 == 0 {
			// front
			newArr[front] = a
			front++
		} else {
			// rear
			newArr[rear] = a
			rear--
		}
	}
	res := abs(newArr[0] - newArr[n-1])
	for i := 0; i < n-1; i++ {
		res = max(res, abs(newArr[i]-newArr[i+1]))
	}

	buf.WriteString(fmt.Sprintf("%d\n", res))
}

func main() {
	br := bufio.NewReader(os.Stdin)
	fmt.Fscanf(br, "%d\n", &t)

	buf := new(bytes.Buffer)
	for i := 0; i < t; i++ {
		fmt.Fscanf(br, "%d\n", &n)
		arr = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscanf(br, "%d", &arr[j])
		}
		fmt.Fscanln(br)
		solve(buf)
	}
	fmt.Print(buf.String())
	buf = nil
}
