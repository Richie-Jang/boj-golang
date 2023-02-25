package main

import "fmt"

var (
	n  int
	dn [10001]int
)

func main() {
	// init
	dn[2] = 1
	for i := 3; i <= 10000; i++ {
		prev := dn[i-1]
		for {
			vv := prev + (prev * prev)
			if vv <= i {
				prev++
			} else {
				break
			}
		}
		dn[i] = prev - 1
	}

	fmt.Scanln(&n)
	var v int
	for i := 0; i < n; i++ {
		fmt.Scanln(&v)
		fmt.Println(dn[v])
	}
}
