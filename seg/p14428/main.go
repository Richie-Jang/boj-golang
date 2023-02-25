package main

import (
	"bytes"
	"fmt"
	"math"
)

type Order struct {
	inst, i, j int
}

const (
	MAX    = 100_001
	INTMAX = 1 << 31
)

var (
	n      int
	arr    []int
	orders []Order
	tree   []int
)

func testCase() {
	n = 5
	arr = []int{0, 5, 4, 3, 2, 1}
	orders = make([]Order, 6)
	orders[0] = Order{2, 1, 3}
	orders[1] = Order{2, 1, 4}
	orders[2] = Order{1, 5, 3}
	orders[3] = Order{2, 3, 5}
	orders[4] = Order{1, 4, 3}
	orders[5] = Order{2, 3, 5}
}

func minIndex(aIndex, bIndex int) int {

	if aIndex == 0 {
		return bIndex
	} else if bIndex == 0 {
		return aIndex
	}

	v1 := arr[aIndex]
	v2 := arr[bIndex]
	if v1 < v2 {
		return aIndex
	} else if v1 > v2 {
		return bIndex
	} else {
		// v1 == v2
		if aIndex <= bIndex {
			return aIndex
		}
		return bIndex
	}
}

func initTree(left, right, idx int) int {
	if left == right {
		tree[idx] = left
		return left
	}
	mid := (left + right) / 2
	v1 := initTree(left, mid, idx*2)
	v2 := initTree(mid+1, right, idx*2+1)
	nIdx := minIndex(v1, v2)
	tree[idx] = nIdx
	return nIdx
}

func query(left, right, idx, sleft, sright int) int {
	if right < sleft || sright < left {
		return 0
	}
	if left <= sleft && sright <= right {
		return tree[idx]
	}
	mid := (left + right) / 2
	v1 := query(left, mid, idx*2, sleft, sright)
	v2 := query(mid+1, right, idx*2+1, sleft, sright)
	return minIndex(v1, v2)
}

func update(left, right, idx, index, nvalue int) int {
	if index < left || right < index {
		return 0
	}
	if left == right {

		return tree[idx]
	}
	mid := (left + right) / 2
	v1 := update(left, mid, idx*2, index, nvalue)
	v2 := update(mid*2+1, right, idx*2+1, index, nvalue)
	nIdx := minIndex(v1, v2)
	tree[idx] = nIdx
	return tree[idx]
}

func solve() {
	buf := new(bytes.Buffer)
	length := 1 << (int(math.Ceil(math.Log2(float64(n)))) + 1)
	fmt.Println("Lenght", length)
	tree = make([]int, length)
	initTree(1, n, 1)

	for _, order := range orders {
		if order.inst == 1 {
			arr[order.i] = order.j
			update(1, n, 1, order.i, order.j)
		} else {
			r := query(1, n, 1, order.i, order.j)
			buf.WriteString(fmt.Sprintf("%d\n", r))
		}
	}
	fmt.Print(buf.String())
	buf = nil
}

func main() {

	testCase()
	solve()

}
