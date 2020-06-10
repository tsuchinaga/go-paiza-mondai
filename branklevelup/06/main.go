package main

import "fmt"

func main() {
	var s, t string
	_, _ = fmt.Scan(&s, &t)

	var ans int
	for i := 0; i <= len(t)-len(s); i++ {
		if s == t[i:i+len(s)] {
			ans++
		}
	}

	fmt.Println(ans)
}
