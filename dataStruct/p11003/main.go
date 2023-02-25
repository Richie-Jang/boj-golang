package main

import (
	"bufio"
	"fmt"
	"os"
)

type ti2 struct {
	value, idx int
}

var dq []ti2

func main() {
	var n, l, g int
	br := bufio.NewReader(os.Stdin)
	bw := bufio.NewWriter(os.Stdout)

	fmt.Fscan(br, &n, &l)

	checkFirst := func(curIdx int) bool {
		ll := len(dq)
		if ll == 0 {
			return false
		}
		v := dq[0]
		return v.idx <= curIdx-l
	}

	checkTwo := func(curV int) bool {
		ll := len(dq)
		if ll == 0 {
			return false
		}
		v := dq[ll-1]
		return v.value > curV
	}

	for i := 1; i <= n; i++ {
		fmt.Fscan(br, &g)

		for {
			ok := checkFirst(i)
			if ok {
				dq = dq[1:]
			} else {
				break
			}
		}

		for {
			ok := checkTwo(g)
			if ok {
				dq = dq[0 : len(dq)-1]
			} else {
				break
			}
		}

		dq = append(dq, ti2{g, i})
		fmt.Fprintf(bw, "%d ", dq[0].value)
	}
	fmt.Fprintln(bw)
	bw.Flush()

}
