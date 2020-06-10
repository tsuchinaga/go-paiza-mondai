package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, d int
	var s string
	_, _ = fmt.Scan(&n)

	ary := make([]sd, 0)
	kmap := make(map[string]int)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&s, &d)
		if j, ok := kmap[s]; ok {
			ary[j].d += d
		} else {
			kmap[s] = len(ary)
			ary = append(ary, sd{s: s, d: d})
		}
	}

	sort.Slice(ary, func(i, j int) bool {
		return ary[i].d > ary[j].d
	})

	for _, sd := range ary {
		fmt.Printf("%s %d\n", sd.s, sd.d)
	}
}

type sd struct {
	s string
	d int
}
