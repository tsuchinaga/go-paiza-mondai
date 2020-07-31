package main

import (
	"fmt"
)

func main() {
	var y, m, d int
	_, _ = fmt.Scan(&y, &m, &d)

	var day int
	// 日を追加
	day += d

	// 月を追加
	day += (m - 1) * 30

	// 年を追加
	day += y * 12 * 30

	if y > 0 {
		// 前年までの閏年の回数だけ1日ずつ追加
		// 0年は閏年になるので1日追加
		day += 1 + (y-1)/4 - (y-1)/100 + (y-1)/400
	}

	wd := []string{"日", "月", "火", "水", "木", "金", "土"}
	fmt.Printf("%s曜日", wd[(3+day)%7]) // 木曜スタートなので3加算(-1年12月30日が水曜日)
}
