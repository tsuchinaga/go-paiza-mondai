package main

import (
	"fmt"
	"time"
)

// 閏年の誤差を無視したいので400年間で行なう
// 年月の1日を作って、1か月後の前日がその月の最終日で、最終日が土日なら最終営業日は金曜日とする
func main() {
	var total int
	wt := make([]int, 7)
	for y := 2000; y <= 2399; y++ {
		for m := 1; m <= 12; m++ {
			total++
			wt[time.Date(y, time.Month(m), 1, 0, 0, 0, 0, time.Local).AddDate(0, 1, -1).Weekday()]++
		}
	}
	wt[5] += wt[0] + wt[6] // 日曜と土曜は金曜日としてカウント

	var x int
	_, _ = fmt.Scan(&x)
	fmt.Printf("%.6f\n", float64(wt[x])/float64(total))
}
