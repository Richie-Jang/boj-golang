package main

import (
	"bufio"
)

var n, m, k int
var data = []int64{}
var writer *bufio.Writer
var tree = []int64{}

func set_tree(n int, v int64, bit, st, ed int) {
	mid := (st + ed) / 2
	if st == ed {
		tree[bit] = v
	}

	if n <= mid {
		set_tree(n, v, 2*bit, st, mid)
	} else {
		set_tree(n, v, 2*bit+1, mid+1, ed)
	}
	tree[bit] = tree[2*bit] + tree[2*bit+1]
}

func main() {
	/* 	reader := bufio.NewReader(os.Stdin)
	   	writer = bufio.NewWriter(os.Stdout)
	   	fmt.Fscanf(reader, "%d %d %d\n", &n, &m, &k)
	   	for i := 1; i <= n; i++ {
	   		v := 0
	   		fmt.Fscanf(reader, "%d\n", &v)
	   		data = append(data, int64(v))
	   	}
	   	for i := 1; i <= m+k; i++ {
	   		a, b, c := 0, 0, 0
	   		fmt.Fscanf(reader, "%d %d %d\n", &a, &b, &c)
	   		conds = append(conds, cond{a, b, int64(c)})
	   	}

	   	solve()
	   	writer.Flush() */
}
