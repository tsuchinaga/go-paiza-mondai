package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	_, _ = fmt.Scan(&s)

	abc := strings.Split(s, "/")
	var y, m, d string

	// abcのうち年に該当するのがなければ no answer
	for i, a := range abc {
		if isYear(a) {
			y = a
			abc = append(abc[:i], abc[i+1:]...)
			break
		}
	}
	if y == "" {
		fmt.Println("no answer")
		return
	}

	// 残った2つのうち、両方ともが月も日も満たすならmany answers
	// またはどちらかが月も日も満たさなければno answer
	am := isMonth(abc[0])
	bm := isMonth(abc[1])
	ad := isDay(abc[0])
	bd := isDay(abc[1])
	if am && bm && ad && bd {
		fmt.Println("many answers")
		return
	} else if (!am && !ad) || (!bm && !bd) {
		fmt.Println("no answer")
		return
	}

	// ここまでこればどちらかが月でどちらかが日であることが確定している
	switch {
	case am && bd:
		m = abc[0]
		d = abc[1]
	case bm && ad:
		m = abc[1]
		d = abc[0]
	}
	fmt.Println(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(s, y, "YYYY"), m, "MM"), d, "DD"))
}

func isYear(a string) bool {
	return len(a) == 4
}

func isMonth(a string) bool {
	return len(a) == 2 && "01" <= a && a <= "12"
}

func isDay(a string) bool {
	return len(a) == 2 && "01" <= a && a <= "31"
}
