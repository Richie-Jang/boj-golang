package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
)

var (
    n,k int
)

type pair struct {
    index, value int
}

func main() {
    br := bufio.NewReader(os.Stdin)
    fmt.Fscanf(br, "%d %d\n", &n, &k)
    arr := make([]int, n)
    for i := 0; i<n; i++ {
        fmt.Fscanf(br, "%d", &arr[i])
    }
    fmt.Fscanln(br)
    sort.Ints(arr)
    total := arr[n-1] - arr[0]

    difArr := make([]pair, 0, n)
    for i := 0; i<n-1; i++ {
        j := i+1
        dif := arr[j]-arr[i]
        difArr = append(difArr, pair{i, dif})
    }

    sort.Slice(difArr, func(i,j int) bool {
        a := difArr[i]
        b := difArr[j]
        if a.value == b.value {
            return a.index > b.index
        }
        return a.value > b.value
    })

    remain := k-1
    idx := 0
    for remain > 0 {
        total -= difArr[idx].value
        idx++
        remain--
    }
    fmt.Println(total)
}
