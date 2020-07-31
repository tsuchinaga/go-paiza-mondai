package main

import (
	"fmt"
	"time"
)

func main() {
	var holiday = map[string]bool{}
	var n, m, d int
	_, _ = fmt.Scan(&n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&m, &d)
		holiday[fmt.Sprintf("%02d%02d", m, d)] = true
	}

	date := time.Date(2018, 12, 31, 0, 0, 0, 0, time.Local)
	var max, h int // 最大日数, 連休数
	for {
		date = date.AddDate(0, 0, 1)
		if !date.Before(time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)) {
			break
		}

		// 日曜日か土曜日か祝日ならcontinue
		if date.Weekday() == 0 || date.Weekday() == 6 || holiday[fmt.Sprintf("%02d%02d", date.Month(), date.Day())] {
			h++
			continue
		}

		if max < h {
			max = h
		}
		h = 0
	}

	if max < h {
		max = h
	}
	fmt.Println(max)
}
