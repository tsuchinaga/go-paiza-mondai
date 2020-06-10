package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, d int
	var s string
	_, _ = fmt.Scan(&n)

	ary := make([]sd, n)
	for i := range ary {
		_, _ = fmt.Scan(&s, &d)
		ary[i] = sd{s: s, d: d}
	}
	sort.Slice(ary, func(i, j int) bool {
		return ary[i].d < ary[j].d
	})
	for _, sd := range ary {
		fmt.Println(sd.s)
	}
}

type sd struct {
	s string
	d int
}
