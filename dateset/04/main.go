package main

import "fmt"

func main() {
	var y, m, day int

	_, _ = fmt.Scan(&y, &m)
	switch m {
	case 2:
		day = 28
		if y%4 == 0 && (y%100 != 0 || y%400 == 0) {
			day++
		}
	case 4, 6, 9, 11:
		day = 30
	default:
		day = 31
	}

	fmt.Println(day)
}
