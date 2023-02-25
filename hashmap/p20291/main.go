package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var (
	n    int
	dict map[string]int
)

func main() {
	br := bufio.NewReader(os.Stdin)
	fmt.Fscanf(br, "%d\n", &n)
	dict = make(map[string]int)
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscanf(br, "%s\n", &s)
		dv := strings.Split(s, ".")
		if a, ok := dict[dv[1]]; ok {
			dict[dv[1]] = a + 1
		} else {
			dict[dv[1]] = 1
		}
	}

	keys := make([]string, 0, len(dict))
	for k := range dict {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, v := range keys {
		fmt.Println(v, dict[v])
	}

}
