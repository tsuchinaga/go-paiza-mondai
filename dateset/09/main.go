package main

import (
	"fmt"
	"time"
)

func main() {
	var m, d int
	var wd string
	_, _ = fmt.Scan(&m, &d, &wd)

	date := time.Date(2001, time.Month(m), d, 0, 0, 0, 0, time.Local)
	wdI := map[string]int{"SUN": 0, "MON": 1, "TUE": 2, "WED": 3, "THU": 4, "FRI": 5, "SAT": 6}[wd]

	date = date.AddDate(0, 0, 1)
	switch wdI {
	case 6:
		date = date.AddDate(0, 0, 1)
	case 5:
		date = date.AddDate(0, 0, 2)
	}

	fmt.Printf("%d月%d日\n", date.Month(), date.Day())
}
