package main

import (
	"fmt"
)

func main() {
	type cd struct {
		b int // 0: 営業日, 1: 休業日
		n int // 日数
	}

	days := make([]cd, 0)
	var n, l, b int
	_, _ = fmt.Scan(&n, &l)

	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&b)
		if len(days) == 0 || days[len(days)-1].b != b {
			days = append(days, cd{n: 1, b: b})
		} else {
			days[len(days)-1].n++
		}
	}

	var max int
	for i := range days {
		var cnt int // 連休数
		h := l      // 休める日数
		for j := i; j < len(days); j++ {
			if days[j].b == 1 { // 休み
				cnt += days[j].n
			} else { // 営業日
				if h < days[j].n { // 有給のほうが少ない
					cnt += h
					h = 0
					break
				} else { // 充分な有給がある
					cnt += days[j].n
					h -= days[j].n
				}
			}
		}
		if max < cnt {
			max = cnt
		}
	}
	fmt.Println(max)
}
