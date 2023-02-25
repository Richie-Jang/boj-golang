package main

import (
	"fmt"
	"math"
)

func main() {

	var t int
	fmt.Scanln(&t)

	var res = float64(t) / 2.0
	res = res * res
	fmt.Println(int(math.Round(res)))
}
