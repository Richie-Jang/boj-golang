package main

import (
	"bufio"
	"fmt"
	"os"
)

func diffSetList(a, b map[int]bool) []int {
	res := make([]int, 0)

	for k := range a {
		_, ok := b[k]
		if !ok {
			res = append(res, k)
		}
	}

	return res
}

func main() {
	var n, m int
	br := bufio.NewReader(os.Stdin)
	fmt.Fscan(br, &n, &m)

	setA := make(map[int]bool)
	setB := make(map[int]bool)
	var g int

	for i := 0; i < n; i++ {
		fmt.Fscan(br, &g)
		setA[g] = true
	}

	for i := 0; i < m; i++ {
		fmt.Fscan(br, &g)
		setB[g] = true
	}

	list1 := diffSetList(setA, setB)
	list2 := diffSetList(setB, setA)
	fmt.Println(len(list1) + len(list2))

}
