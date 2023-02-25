package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	bf := bufio.NewReader(os.Stdin)
	var a, b, c int
	fmt.Fscanf(bf, "%d %d %d\n", &a, &b, &c)

	if b >= c {
		fmt.Println("-1")
	} else {
		res := a/(c-b) + 1
		fmt.Println(res)
	}

}
