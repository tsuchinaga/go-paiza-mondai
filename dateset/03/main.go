package main

import (
	"fmt"
	"time"
)

func main() {
	var y, m, d, x int
	var g string

	_, _ = fmt.Scan(&y, &m, &d)
	day := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local)

	if !day.Before(time.Date(2019, 5, 1, 0, 0, 0, 0, time.Local)) {
		g = "令和"
		x = y - 2019
	} else if !day.Before(time.Date(1989, 1, 8, 0, 0, 0, 0, time.Local)) {
		g = "平成"
		x = y - 1989
	} else if !day.Before(time.Date(1926, 12, 25, 0, 0, 0, 0, time.Local)) {
		g = "昭和"
		x = y - 1926
	} else if !day.Before(time.Date(1912, 7, 30, 0, 0, 0, 0, time.Local)) {
		g = "大正"
		x = y - 1912
	} else {
		g = "明治"
		x = y - 1868
	}

	if x == 0 {
		fmt.Printf("%s元年%d月%d日\n", g, m, d)
	} else {
		fmt.Printf("%s%d年%d月%d日\n", g, x+1, m, d)
	}
}
