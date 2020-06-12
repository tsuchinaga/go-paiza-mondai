package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	if n >= 80 {
		fmt.Println("pass")
	} else {
		fmt.Println("fail")
	}
}
