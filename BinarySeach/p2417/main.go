package main

import (
	"fmt"
	"math"
)

func main() {
	var n uint64
	fmt.Scanln(&n)

	v := math.Sqrt(float64(n))
	v = math.Ceil(v)
	fmt.Println(uint64(v))
}
