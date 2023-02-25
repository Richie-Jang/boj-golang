package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct {
	starter int
	ender   int
}

var n, m, k int
var data []int
var tree = make(map[pair]uint64)

const DIV = 1000000007

func setupTree(st, et int) uint64 {
	if st == et {
		tree[pair{st, st}] = uint64(data[st])
		return uint64(data[st])
	}
	mid := (st + et) / 2
	left := setupTree(st, mid)
	right := setupTree(mid+1, et)
	sum := (left * right) % DIV
	tree[pair{st, et}] = sum
	return sum
}

func updateValue(ind, val, st, et int) uint64 {
	if !(st <= ind && ind <= et) {
		return tree[pair{st, et}]
	}
	if st == et {
		tree[pair{st, et}] = uint64(val)
		return uint64(val)
	}
	m := (st + et) / 2
	l := updateValue(ind, val, st, m)
	r := updateValue(ind, val, m+1, et)
	rr := (l * r) % DIV
	tree[pair{st, et}] = rr
	return rr
}

func printValue(ind1, ind2, st, et int, ans *int) {
	if ind1 > et || ind2 < st {
		return
	}
	if ind1 <= st && et <= ind2 {
		*ans = int((uint64(*ans) * tree[pair{st, et}]) % DIV)
		return
	}

	m := (st + et) / 2
	printValue(ind1, ind2, st, m, ans)
	printValue(ind1, ind2, m+1, et, ans)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n, &m, &k)
	data = make([]int, n+1)
	var v int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &v)
		data[i+1] = v
	}

	setupTree(1, n)

	var v2, v3 int
	for i := 0; i < (m + k); i++ {
		fmt.Fscanln(reader, &v, &v2, &v3)
		switch v {
		case 1:
			updateValue(v2, v3, 1, n)
			//fmt.Println(tree)
		case 2:
			var ans = 1
			printValue(v2, v3, 1, n, &ans)
			fmt.Println(ans)
		}
	}

}
