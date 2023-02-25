package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Dr struct {
	from, to, weight int
}

var n, m int
var adjs map[int][]Dr

func getDist(a, b int) int {
	vs := make([]int, n+1)
	dst := make([]int, 0)
	var rec func(c, acc int)

	rec = func(c, acc int) {
		if c == b {
			dst = append(dst, acc)
			return
		}
		ns, ex := adjs[c]
		vs[c] = 1
		if ex {
			for _, v := range ns {
				v2 := v.to
				if vs[v2] == 1 {
					continue
				}
				rec(v2, acc+v.weight)
			}
		}
	}

	rec(a, 0)

	if len(dst) == 0 {
		return 0
	}
	if len(dst) == 1 {
		return dst[0]
	}
	sort.Ints(dst)
	return dst[0]
}

func main() {
	br := bufio.NewReader(os.Stdin)
	bw := bufio.NewWriter(os.Stdout)
	adjs = make(map[int][]Dr)
	fmt.Fscan(br, &n, &m)
	a, b, c := 0, 0, 0
	for i := 0; i < n-1; i++ {
		fmt.Fscan(br, &a, &b, &c)
		if _, ex := adjs[a]; ex {
			adjs[a] = append(adjs[a], Dr{a, b, c})
		} else {
			adjs[a] = []Dr{{a, b, c}}
		}
		if _, ex := adjs[b]; ex {
			adjs[b] = append(adjs[b], Dr{b, a, c})
		} else {
			adjs[b] = []Dr{{b, a, c}}
		}
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(br, &a, &b)
		fmt.Fprintf(bw, "%d\n", getDist(a, b))
	}
	bw.Flush()
}
