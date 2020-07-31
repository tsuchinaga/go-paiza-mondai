package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var y, m int
	_, _ = fmt.Scan(&y, &m)
	date := time.Date(y, time.Month(m), 1, 0, 0, 0, 0, time.Local)
	y, x := 0, int(date.Weekday())
	cal := [6][7]string{}

	for int(date.Month()) == m {
		cal[y][x] = fmt.Sprintf("%2d", date.Day())
		y, x = y+(x+1)/7, (x+1)%7
		date = date.AddDate(0, 0, 1)
	}

	for i := range cal {
		for j := range cal[i] {
			if cal[i][j] == "" {
				cal[i][j] = "  "
			}
		}
		fmt.Println(strings.Join(cal[i][:], " "))
	}
}
