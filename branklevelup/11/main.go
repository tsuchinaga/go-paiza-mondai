package main

import "fmt"

func main() {
	var n, m int
	var name, blood, res string

	_, _ = fmt.Scan(&n)
	names := make([]string, n)
	nb := make(map[string]string)
	br := make(map[string]string)

	for i := range names {
		_, _ = fmt.Scan(&name, &blood)
		names[i] = name
		nb[name] = blood
	}

	_, _ = fmt.Scan(&m)
	for i := 0; i < m; i++ {
		_, _ = fmt.Scan(&blood, &res)
		br[blood] = res
	}

	for _, name = range names {
		fmt.Printf("%s %s\n", name, br[nb[name]])
	}
}
