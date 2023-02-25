package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func adjustSlice(maxLen int, arr []int, v string) {
	curLen := len(v)
	point_maxLen := maxLen
	cur := 0
	subLen := 0
	for cur < maxLen {
		if curLen < point_maxLen {
			point_maxLen--
			subLen++
			arr[cur] = 0
		} else {
			arr[cur] = int(v[cur-subLen] - '0')
		}
		cur++
	}
}

func main() {
	var a, b string
	fmt.Scanln(&a, &b)
	alen := len(a)
	blen := len(b)

	ll := alen
	if alen < blen {
		ll = blen
	}

	a1 := make([]int, ll)
	b1 := make([]int, ll)

	adjustSlice(ll, a1, a)
	adjustSlice(ll, b1, b)

	res := make([]int, ll)

	for i := 0; i < ll; i++ {
		aa := a1[i]
		bb := b1[i]
		res[i] = aa + bb
	}

	var buf bytes.Buffer
	for i := 0; i < ll; i++ {
		buf.WriteString(strconv.FormatInt(int64(res[i]), 10))
	}

	fmt.Println(buf.String())

}
