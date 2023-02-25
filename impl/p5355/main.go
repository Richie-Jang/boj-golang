package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var n int

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ = strconv.Atoi(sc.Text())
	for i := 0; i < n; i++ {
		sc.Scan()
		ss := strings.Split(sc.Text(), " ")
		v, _ := strconv.ParseFloat(ss[0], 32)
		for j := 1; j < len(ss); j++ {
			switch ss[j] {
			case "@":
				v *= 3.0
			case "%":
				v += 5.0
			case "#":
				v -= 7.0
			}
		}
		fmt.Printf("%.2f\n", v)
	}
}
