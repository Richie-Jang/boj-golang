package main

import "fmt"

var sum int

func main() {
	for {
		var v int
		_, err := fmt.Scanln(&v)
		if err != nil {
			break
		}
		sum += v
	}

	fmt.Println(sum)
}
