package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	n int
)

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	br := bufio.NewScanner(os.Stdin)
	br.Scan()
	n, _ = strconv.Atoi(br.Text())
	var line string
	buf := new(bytes.Buffer)
	for i := 0; i < n; i++ {
		br.Scan()
		line = br.Text()
		arr := strings.Split(line, " ")
		rrr := make([]int, 0, len(arr))
		for j := 0; j < len(arr)-1; j++ {
			a1, _ := strconv.Atoi(arr[j])
			for k := j + 1; k < len(arr); k++ {
				a2, _ := strconv.Atoi(arr[k])
				rrr = append(rrr, gcd(a1, a2))
			}
		}
		sort.Ints(rrr)
		buf.WriteString(fmt.Sprintf("%d\n", rrr[len(rrr)-1]))
	}

	fmt.Print(buf.String())

}
