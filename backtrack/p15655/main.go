package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var n, m int
var arr []int

var bt func(idx, k int, acc []int)

func solve() {
	sort.Ints(arr)
	res := make([][]int, 0, n)

	bt = func(idx, k int, acc []int) {
		if k == 0 {
			rr := make([]int, len(acc))
			copy(rr, acc)
			res = append(res, rr)
			return
		}

		for i := idx; i < n; i++ {
			acc = append(acc, arr[i])
			bt(i+1, k-1, acc)
			acc = acc[:len(acc)-1]
		}

	}

	bt(0, m, make([]int, 0))
	// print
	var buf bytes.Buffer
	for _, v := range res {
		rr := make([]string, len(v))
		for j := 0; j < len(v); j++ {
			rr[j] = strconv.Itoa(v[j])
		}
		buf.WriteString(strings.Join(rr, " "))
		buf.WriteString("\n")
	}
	fmt.Print(buf.String())
}

func main() {

	br := bufio.NewReader(os.Stdin)
	fmt.Fscanf(br, "%d %d\n", &n, &m)
	arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(br, "%d", &arr[i])
	}
	fmt.Fscanln(br)
	solve()

}
