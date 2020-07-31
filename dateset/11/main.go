package main

import (
	"fmt"
	"time"
)

func main() {
	var date, before, after time.Time
	date = time.Date(2018, 12, 31, 0, 0, 0, 0, time.Local)
	before, after = date.AddDate(0, 0, -1), date.AddDate(0, 0, 1)

	var max, n int // 最大日数, 連休数
	for date.Before(time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)) {
		before, date, after = date, after, after.AddDate(0, 0, 1)
		w := date.Weekday()

		// 日曜日か土曜日か祝日ならcontinue
		if w == 0 || w == 6 || holiday[fmt.Sprintf("%02d%02d", date.Month(), date.Day())] {
			n++
			continue
		}

		// 前日が祝日で日曜日ならcontinue
		if w-1 == 0 && holiday[fmt.Sprintf("%02d%02d", before.Month(), before.Day())] {
			n++
			continue
		}

		// 前日と翌日が祝日ならcontinue
		if holiday[fmt.Sprintf("%02d%02d", before.Month(), before.Day())] && holiday[fmt.Sprintf("%02d%02d", after.Month(), after.Day())] {
			n++
			continue
		}

		if max < n {
			max = n
		}
		n = 0
	}

	fmt.Println(max)
}

var holiday = map[string]bool{
	"0101": true,
	"0114": true,
	"0211": true,
	"0321": true,
	"0429": true,
	"0501": true,
	"0503": true,
	"0504": true,
	"0505": true,
	"0715": true,
	"0811": true,
	"0916": true,
	"0923": true,
	"1014": true,
	"1022": true,
	"1103": true,
	"1123": true,
}
