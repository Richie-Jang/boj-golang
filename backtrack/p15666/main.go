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

var (
	n, m  int
	items map[int]any
)

func solve() {
	arr := make([]int, len(items))
	idx := 0
	for k, _ := range items {
		arr[idx] = k
		idx++
	}
	sort.Ints(arr)

	var loop func(int, int, []int)
	res := make([][]int, 0)

	loop = func(index, count int, acc []int) {
		if count == 0 {
			vv := make([]int, len(acc))
			copy(vv, acc)
			res = append(res, vv)
			return
		}
		for v := index; v < len(arr); v++ {
			acc = append(acc, arr[v])
			loop(v, count-1, acc)
			acc = acc[:len(acc)-1]
		}
	}

	loop(0, m, make([]int, 0))

	var buf bytes.Buffer
	for _, v := range res {
		sv := make([]string, len(v))
		for i, b := range v {
			sv[i] = strconv.Itoa(b)
		}
		buf.WriteString(strings.Join(sv, " "))
		buf.WriteString("\n")
	}

	fmt.Print(buf.String())
}

func main() {
	br := bufio.NewReader(os.Stdin)
	fmt.Fscanf(br, "%d %d\n", &n, &m)
	items = make(map[int]any)
	var v int
	for i := 0; i < n; i++ {
		fmt.Fscanf(br, "%d", &v)
		items[v] = nil
	}
	fmt.Fscanln(br)

	solve()
}
