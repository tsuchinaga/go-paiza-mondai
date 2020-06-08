package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, g, s int
	_, _ = fmt.Scan(&n)

	ps := make([]owned, n)
	for i := range ps {
		_, _ = fmt.Scan(&g, &s)
		ps[i] = owned{g: g, s: s}
	}

	sort.Slice(ps, func(i, j int) bool {
		return ps[i].s > ps[j].s || (ps[i].s == ps[j].s && ps[i].g > ps[j].g)
	})

	for _, p := range ps {
		fmt.Printf("%d %d\n", p.g, p.s)
	}
}

type owned struct {
	g int
	s int
}
