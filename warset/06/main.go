package main

import (
	"fmt"
)

func main() {
	var s, t string
	_, _ = fmt.Scan(&s, &t)
	fmt.Printf("%s.%s.", string(s[0]), string(t[0]))
}
