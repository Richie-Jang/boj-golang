package main

import "fmt"

var n, m string

func main() {
	fmt.Scanln(&n, &m)
	var sum int
	for i := 0; i < len(n); i++ {
		for j := 0; j < len(m); j++ {
			d1 := int(n[i] - '0')
			d2 := int(m[j] - '0')
			sum += (d1 * d2)
		}
	}
	fmt.Println(sum)
}
