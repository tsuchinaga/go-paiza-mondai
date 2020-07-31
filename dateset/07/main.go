package main

import (
	"fmt"
	"time"
)

func main() {
	var y, m, d int
	_, _ = fmt.Scan(&y, &m, &d)
	date := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local)

	wd := []string{"日", "月", "火", "水", "木", "金", "土"}
	fmt.Printf("%s曜日", wd[date.Weekday()])
}
