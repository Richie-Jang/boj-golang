package main

import "fmt"

var n, a0, a1 int

func main() {

	fmt.Scan(&n)
	count := 2013
	a0, a1 = 9, 5
	sign := 1
	if n < count {
		sign = -1
	}
	for {
		if count == n {
			break
		}
		count += sign
		a0 = a0 + sign
		a1 = a1 + sign
		// correction
		if a0 < 0 {
			a0 = 9
		} else if a0 >= 10 {
			a0 = 0
		}
		if a1 < 0 {
			a1 = 11
		} else if a1 >= 12 {
			a1 = 0
		}

	}

	fmt.Printf("%c%d\n", rune('A')+rune(a1), a0)

}
