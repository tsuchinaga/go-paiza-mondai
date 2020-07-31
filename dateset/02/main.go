package main

import (
	"fmt"
	"time"
)

func main() {
	var y, m, d int
	_, _ = fmt.Scan(&y, &m, &d)
	day := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local)

	var g string
	if day.Before(time.Date(1912, 7, 30, 0, 0, 0, 0, time.Local)) {
		g = "明治"
	} else if day.Before(time.Date(1926, 12, 25, 0, 0, 0, 0, time.Local)) {
		g = "大正"
	} else if day.Before(time.Date(1989, 1, 8, 0, 0, 0, 0, time.Local)) {
		g = "昭和"
	} else if day.Before(time.Date(2019, 5, 1, 0, 0, 0, 0, time.Local)) {
		g = "平成"
	} else {
		g = "令和"
	}
	fmt.Printf("%s年%d月%d日\n", g, m, d)
}
