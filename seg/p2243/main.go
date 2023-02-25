package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
)

type Order struct {
	tp, taste, addNum int
}

const (
	MAX = 1_000_000
)

var (
	n      int
	orders []Order
	tree   []int
	buf    bytes.Buffer
)

func queryIndex(left, right, nodeIdx int, target int) int {
	if left == right {
		return left
	}
	mid := (left + right) / 2
	if target <= tree[nodeIdx*2] {
		return queryIndex(left, mid, nodeIdx*2, target)
	} else {
		return queryIndex(mid+1, right, nodeIdx*2+1, target-tree[nodeIdx*2])
	}
}

func update(left, right, nodeIdx, target int, diff int) {
	if target < left || right < target {
		return
	}

	tree[nodeIdx] += diff
	if left != right {
		mid := (left + right) / 2
		update(left, mid, nodeIdx*2, target, diff)
		update(mid+1, right, nodeIdx*2+1, target, diff)
	}
}

func solve() {
	segHeight := int(math.Ceil(math.Log2(MAX)))
	segSize := 1 << (segHeight + 1)
	tree = make([]int, segSize)

	for _, order := range orders {
		switch order.tp {
		case 1:
			idx := queryIndex(1, MAX, 1, order.taste)
			buf.WriteString(fmt.Sprintf("%d\n", idx))
			update(1, MAX, 1, idx, -1)
		default:
			update(1, MAX, 1, order.taste, order.addNum)
		}
	}
	fmt.Print(buf.String())
}

func main() {
	br := bufio.NewReader(os.Stdin)
	fmt.Fscanf(br, "%d\n", &n)
	var c, d, e int
	orders = make([]Order, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(br, "%d", &c)
		if c == 1 {
			fmt.Fscanf(br, "%d\n", &d)
			e = 0
		} else {
			fmt.Fscanf(br, "%d %d\n", &d, &e)
		}
		orders[i] = Order{c, d, e}
	}

	solve()
}
