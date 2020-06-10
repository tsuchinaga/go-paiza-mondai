package main

import "fmt"

func main() {
	var c, s string
	_, _ = fmt.Scan(&c, &s)

	var ans int
	for _, r := range []rune(s) {
		if string(r) == c {
			ans++
		}
	}

	fmt.Println(ans)
}
