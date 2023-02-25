package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)

	fmt.Scanln(a)
	fmt.Scanln(b)

	c := new(big.Int)
	d := new(big.Int)
	e := new(big.Int)
	fmt.Println(c.Add(a, b))
	fmt.Println(d.Sub(a, b))
	fmt.Println(e.Mul(a, b))
}
