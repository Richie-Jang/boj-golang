package main

import (
	"bufio"
	"fmt"
	"os"
)

var gates []int

func findGate(g int) int {
	if g == gates[g] {
		return g
	}
	gates[g] = findGate(gates[g])
	return gates[g]
}

func union(g, gi int) {
	g = findGate(g)
	gi = findGate(gi)

	if g != gi {
		gates[g] = gi
	}
}

func main() {
	var g, f int
	br := bufio.NewReader(os.Stdin)
	fmt.Fscan(br, &g)
	fmt.Fscan(br, &f)
	gates = make([]int, g+1)
	for i := range gates {
		gates[i] = i
	}
	var gi, count int
	isFinish := false
	for i := 1; i <= f; i++ {
		fmt.Fscan(br, &gi)
		if isFinish {
			continue
		}
		eGate := findGate(gi)
		if eGate == 0 {
			isFinish = true
			continue
		}

		count++
		union(eGate, eGate-1)

	}

	fmt.Println(count)
}
