package main

import (
	"fmt"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)

	var c string

	names := make(map[string]int)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&c)
		if _, ok := names[c]; !ok {
			names[c] = 1
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
