package main

import "fmt"

func main() {
	var n, a, ans int
	_, _ = fmt.Scan(&n)

	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&a)
		if a >= 5 {
			ans += a
		}
	}

	fmt.Println(ans)
}
