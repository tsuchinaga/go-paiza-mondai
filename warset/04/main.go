package main

import (
	"fmt"
)

func main() {
	var n int
	var d string
	_, _ = fmt.Scan(&n)

	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&d)
		switch d {
		case "forward":
			fmt.Println("Sunny")
		case "reverse":
			fmt.Println("Rainy")
		case "sideways":
			fmt.Println("Cloudy")
		}
	}
}
