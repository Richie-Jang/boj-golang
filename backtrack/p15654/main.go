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

var n,m int
var arr []int
var btfunc func(k int, set map[int]any, acc []int)
func solve() {
    res := make([][]int, 0, n)
    btfunc = func(k int, set map[int]any, acc []int) {
        if k == 0 {
            rr := make([]int, len(acc))
            copy(rr, acc)
            res = append(res, rr)
            return
        }

        for _,v := range arr {
            if _, ok := set[v]; !ok {
                set[v] = nil
                acc = append(acc, v)
                btfunc(k-1, set, acc)
                delete(set, v)
                acc = acc[:len(acc)-1]
            }
        }

    }

    btfunc(m, make(map[int]any), make([]int, 0))

    // output
    var buf bytes.Buffer
    for _, r := range res {
        rstr := make([]string, len(r))
        for i := 0; i < len(r); i++ {
            rstr[i] = strconv.Itoa(r[i])
        }
        buf.WriteString(strings.Join(rstr, " "))
        buf.WriteString("\n")
        rstr = nil
    }
    fmt.Println(buf.String())
}

func main() {
    br := bufio.NewReader(os.Stdin)
    fmt.Fscanf(br, "%d %d\n", &n, &m)
    arr = make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Fscanf(br, "%d", &arr[i])
    }
    fmt.Fscanln(br)
    sort.Ints(arr)
    solve()
}