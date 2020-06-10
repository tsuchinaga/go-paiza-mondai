package main

import (
	"fmt"
)

func main() {
	var x, y, c string
	_, _ = fmt.Scan(&x, &y, &c)
	if x <= c && c <= y {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
