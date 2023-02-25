package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func countChar(s string) string {
	var a1, a2, a3, a4 int

	checkBigChars := func(c rune) bool {
		st, ed := 'A', 'Z'
		return c >= st && c <= ed
	}

	checkSmallChars := func(c rune) bool {
		st, ed := 'a', 'z'
		return c >= st && c <= ed
	}

	for _, c := range s {
		if c == ' ' {
			a4++
		} else if checkBigChars(c) {
			a2++
		} else if checkSmallChars(c) {
			a1++
		} else {
			a3++
		}
	}

	return fmt.Sprintf("%d %d %d %d\n", a1, a2, a3, a4)
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	bw := new(bytes.Buffer)

	for {
		l, _, err := rd.ReadLine()
		if err != nil {
			break
		}
		ll := string(l)
		bw.WriteString(countChar(ll))
	}

	fmt.Print(bw.String())
}
