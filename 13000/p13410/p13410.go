package main

import (
	"fmt"
	"math"
)

var (
	n, k int
)

func reverse(v int) int {
	if v < 10 {
		return v
	}

	vl := int(math.Log10(float64(v)))
	g := v
	acc := 0
	for i := vl; i >= 0; i-- {
		p := int(math.Pow10(i))
		m := g / p
		g = g - (m * p)
		vv := vl - i
		p2 := int(math.Pow10(vv))
		acc += p2 * m
	}
	return acc

}

func main() {
	fmt.Scanf("%d %d\n", &n, &k)
	res := 0
	for i := 1; i <= k; i++ {
		v := n * i
		//fmt.Println(n, "*", i, "=", v, "=>", reverse(v))
		v2 := reverse(v)
		if res < v2 {
			res = v2
		}
	}

	fmt.Println(res)
}
