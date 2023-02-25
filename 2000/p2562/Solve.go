package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	bf := bufio.NewReader(os.Stdin)
	var v int
	var bigValue int
	var bigIndex int
	for lineCount := 1; lineCount <= 9; lineCount++ {

		fmt.Fscanf(bf, "%d\n", &v)
		if bigValue < v {
			bigValue = v
			bigIndex = lineCount
		}

	}
	fmt.Println(bigValue)
	fmt.Println(bigIndex)
}
