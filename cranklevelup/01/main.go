package main

import "fmt"

func main() {
	var n, a int
	var s string

	_, _ = fmt.Scan(&n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&s, &a)
		fmt.Printf("%s %d\n", s, a+1)
	}
}
