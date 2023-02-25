package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type exRange struct {
	start, end int
}

var (
	n, t                    int
	preOrd, inOrd           []int
	preOrdIndex, inOrdIndex []int
	buf                     *bytes.Buffer
)

func getMinMaxByOrders(rg exRange) (int, int) {
	min, max := n, 0
	for i := rg.start; i < rg.end; i++ {
		v := inOrd[i]
		v2 := preOrdIndex[v]
		if min > v2 {
			min = v2
		}
		if max < v2 {
			max = v2
		}
	}
	return min, max
}

func solve(od1, od2 exRange, postOrd *[]int) {

	//fmt.Println(od1, od2, *postOrd)
	if len(*postOrd) == n {
		for i := n - 1; i >= 0; i-- {
			buf.WriteString(fmt.Sprintf("%d ", (*postOrd)[i]))
		}
		buf.WriteString("\n")
		return
	}

	if od1.start == -1 {
		return
	}

	root := preOrd[od1.start]
	*postOrd = append(*postOrd, root)
	// right first
	inIndex := inOrdIndex[root]
	//fmt.Println("found index: ", inIndex, "for Root:", root)
	r2 := exRange{inIndex + 1, od2.end}

	if od1.end-od1.start == 1 {
		solve(exRange{-1, -1}, exRange{-1, -1}, postOrd)
	}

	if r2.end-r2.start > 0 {
		m1, m2 := getMinMaxByOrders(r2)
		//fmt.Println("right : min-max", m1, m2)
		rightSide := exRange{m1, m2 + 1}
		if rightSide.end-rightSide.start > 0 {
			solve(rightSide, r2, postOrd)
		}
	}

	// left
	l2 := exRange{od2.start, inIndex}
	if l2.end-l2.start > 0 {
		m1, m2 := getMinMaxByOrders(l2)
		//fmt.Println("left : min-max", m1, m2)
		leftSide := exRange{m1, m2 + 1}
		if leftSide.end-leftSide.start > 0 {
			solve(leftSide, l2, postOrd)
		}
	}

}

func main() {
	br := bufio.NewReader(os.Stdin)
	buf = new(bytes.Buffer)
	fmt.Fscanf(br, "%d\n", &t)
	for t > 0 {
		fmt.Fscanf(br, "%d\n", &n)
		preOrd = make([]int, n)
		inOrd = make([]int, n)

		preOrdIndex = make([]int, n+1)
		inOrdIndex = make([]int, n+1)

		for i := 0; i < n; i++ {
			fmt.Fscanf(br, "%d", &preOrd[i])
			preOrdIndex[preOrd[i]] = i
		}
		fmt.Fscanln(br)
		for i := 0; i < n; i++ {
			fmt.Fscanf(br, "%d", &inOrd[i])
			inOrdIndex[inOrd[i]] = i
		}
		fmt.Fscanln(br)

		postOrd := make([]int, 0, n)
		solve(exRange{0, n}, exRange{0, n}, &postOrd)
		t--
	}

	fmt.Print(buf.String())
}
