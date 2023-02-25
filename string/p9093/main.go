package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

var n int

func reverseStr(s string) string {
	if len(s) == 1 {
		return s
	}

	res := make([]byte, len(s))
	st, ed := 0, len(s)-1

	for {
		if st > ed {
			break
		} else if st == ed {
			res[st] = s[st]
		} else {
			res[st] = s[ed]
			res[ed] = s[st]
		}
		st++
		ed--
	}
	return string(res)
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	bw := new(bytes.Buffer)

	fmt.Fscanln(rd, &n)

	for i := 0; i < n; i++ {
		l, _, _ := rd.ReadLine()
		//fmt.Println("reading: ", string(l))
		strReader := strings.NewReader(string(l))

		strs := make([]string, 0)
		for {
			str := ""
			_, err := fmt.Fscan(strReader, &str)
			if err != nil {
				break
			}
			strs = append(strs, reverseStr(str))
		}

		bw.WriteString(strings.Join(strs, " ") + "\n")

	}

	fmt.Print(bw.String())
}
