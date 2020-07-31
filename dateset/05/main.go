package main

import (
	"fmt"
	"time"
)

func main() {
	var y, m, d int
	_, _ = fmt.Scan(&y, &m, &d)
	date := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local)
	date = date.AddDate(0, 0, 1)
	fmt.Printf("%d %d %d\n", date.Year(), date.Month(), date.Day())
}
