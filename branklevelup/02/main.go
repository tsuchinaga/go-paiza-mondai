package main

import "fmt"

func main() {
	var n, a, b, ans int
	_, _ = fmt.Scan(&n)

	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&a, &b)
		if a == b {
			ans += a * b
		} else {
			ans += a + b
		}
	}

	fmt.Println(ans)
}
