package main

import (
	"fmt"
)

func main() {
	var n, m, k, a int
	_, _ = fmt.Scan(&n, &m, &k)
	scores := make([]int, n)

	for i := range scores {
		for j := 0; j < m; j++ {
			_, _ = fmt.Scan(&a)
			if k == a {
				scores[i]++
			}
		}
	}

	for _, s := range scores {
		fmt.Println(s)
	}
}
