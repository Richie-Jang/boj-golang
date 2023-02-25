package main

import (
	"bufio"
	"fmt"
	"os"
)

var n int
var people []int
var mm = make(map[int]int)

func solve(sum int) {
	count := n / 2
	for i := 0; i < n; i++ {
		var pSum = 0
		for j := 0; j < count; j++ {
			pos := i + (j * 2)
			if pos >= n {
				pos = pos % n
			}
			pSum += people[pos]
		}
		pp := i - 1
		if pp < 0 {
			pp = n - 1
		}
		mm[pp] = sum - pSum
	}

	for i := 0; i < n; i++ {
		fmt.Println(mm[i])
	}
}

func main() {
	sc := bufio.NewReader(os.Stdin)
	fmt.Fscanf(sc, "%d\n", &n)
	people = make([]int, n)
	var v int
	var sum int
	for i := 0; i < n; i++ {
		fmt.Fscanf(sc, "%d\n", &v)
		people[i] = v
		sum += v
	}

	solve(sum / 2)
}
