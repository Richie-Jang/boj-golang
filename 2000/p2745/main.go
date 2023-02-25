package main

import (
	"fmt"
	"math"
)

var table = make(map[byte]int)

func convertDigit(bs []byte, digit int) int {
	size := len(bs)

	sum := 0
	for i, v := range bs {
		v1 := 0
		if v >= 65 {
			v1 = table[v]
		} else {
			v1 = int(v - '0')
		}
		ii := size - i - 1
		a1 := math.Pow(float64(digit), float64(ii))
		sum += (int(a1) * v1)
	}

	return sum
}

func main() {
	// setup
	vv := 10
	for i := 'A'; i <= 'Z'; i++ {
		table[byte(i)] = vv
		vv++
	}

	var n string
	var b int
	fmt.Scanln(&n, &b)
	fmt.Println(convertDigit([]byte(n), b))
}
