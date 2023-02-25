package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct {
	first, second byte
}

var n int

var (
	smallRange = pair{97, 122}
	bigRange   = pair{65, 90}
)

func main() {
	fmt.Scanln(&n)
	sc := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		sc.Scan()
		str := sc.Text()
		buf := []byte(str)
		c0 := buf[0]
		if c0 >= smallRange.first && c0 <= smallRange.second {
			buf[0] = buf[0] - 32
			fmt.Println(string(buf))
		} else {
			fmt.Println(str)
		}
	}
}
