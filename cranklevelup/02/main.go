package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n, h, m int
	var t string

	_, _ = fmt.Scan(&n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&t, &h, &m)
		var th, tm int
		ts := strings.Split(t, ":")
		th, _ = strconv.Atoi(ts[0])
		tm, _ = strconv.Atoi(ts[1])

		th += h
		tm += m

		th += tm / 60
		tm %= 60
		th %= 24

		fmt.Printf("%02d:%02d\n", th, tm)
	}
}
