package main

import (
	"fmt"
	"time"
)

func main() {
	var m, d int
	var wd string
	_, _ = fmt.Scan(&m, &d, &wd)
	var date, before, after time.Time
	date = time.Date(2019, time.Month(m), d, 0, 0, 0, 0, time.Local)
	before, after = date.AddDate(0, 0, -1), date.AddDate(0, 0, 1)
	w := wdI[wd]

	var n int // 何日先になるか
	for {
		n++
		before, date, after = date, after, after.AddDate(0, 0, 1)
		w = (w + 1) % 7

		// 日曜日か土曜日か祝日ならcontinue
		if w == 0 || w == 6 || holiday[fmt.Sprintf("%02d%02d", date.Month(), date.Day())] {
			continue
		}

		// 前日が祝日で日曜日ならcontinue
		if w-1 == 0 && holiday[fmt.Sprintf("%02d%02d", before.Month(), before.Day())] {
			continue
		}

		// 前日と翌日が祝日ならcontinue
		if holiday[fmt.Sprintf("%02d%02d", before.Month(), before.Day())] && holiday[fmt.Sprintf("%02d%02d", after.Month(), after.Day())] {
			continue
		}

		break
	}

	fmt.Printf("%d月%d日\n", date.Month(), date.Day())
}

var wdI = map[string]int{"SUN": 0, "MON": 1, "TUE": 2, "WED": 3, "THU": 4, "FRI": 5, "SAT": 6}

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
