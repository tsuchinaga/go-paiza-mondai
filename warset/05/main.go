package main

import (
	"fmt"
)

func main() {
	var m int
	_, _ = fmt.Scan(&m)

	if m <= 5 {
		fmt.Println("raw")
	} else if m <= 7 {
		fmt.Println("soft boiled")
	} else if m <= 15 {
		fmt.Println("hard boiled")
	}
}
