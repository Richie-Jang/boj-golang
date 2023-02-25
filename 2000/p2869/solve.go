package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b, c int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d %d\n", &a, &b, &c)

	days := (c - b) / (a - b)
	if (c-b)%(a-b) > 0 {
		days++
	}
	fmt.Println(days)

}
