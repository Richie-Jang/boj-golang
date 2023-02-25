package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	orgs := [6]int{1, 1, 2, 2, 2, 8}
	ps := [6]int{}
	fmt.Scanln(&ps[0], &ps[1], &ps[2], &ps[3], &ps[4], &ps[5])

	nps := [6]string{}

	for i := 0; i < len(orgs); i++ {
		nps[i] = strconv.Itoa(orgs[i] - ps[i])
	}

	fmt.Println(strings.Join(nps[:], " "))

}
