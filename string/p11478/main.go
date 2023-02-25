package main

import (
	"fmt"
)

func main() {

	var inp string
	fmt.Scanf("%s\n", &inp)
	mySet := make(map[string]bool)

	var slice []byte
	for i := 0; i < len(inp); i++ {
		slice = make([]byte, 0)
		for j := i; j < len(inp); j++ {
			slice = append(slice, inp[j])
			mySet[string(slice)] = true
		}
	}

	fmt.Println(len(mySet))
}
