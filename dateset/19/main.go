package main

import (
	"fmt"
	"time"
)

func main() {
	// 閏年の誤差を無視したいので400年間の13日の金曜日の日数と全日数を計算する
	var tf, td int // 13日の金曜の数, 全ての13日の日数
	for y := 2000; y < 2400; y++ {
		for m := 1; m <= 12; m++ {
			if time.Date(y, time.Month(m), 13, 0, 0, 0, 0, time.Local).Weekday() == 5 {
				tf++
			}
		}
		td += 12
	}
	fmt.Printf("%.3f\n", float64(tf)/float64(td))
}
