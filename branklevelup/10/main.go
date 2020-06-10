package main

import (
	"fmt"
	"strings"
)

func main() {
	var line string

	ary := make([]int, 25)
	for i := 0; i < 5; i++ {
		_, _ = fmt.Scan(&line)
		for j, s := range strings.Split(line, "") {
			switch s {
			case "O":
				ary[i*5+j] = 1
			case "X":
				ary[i*5+j] = -1
			}
		}
	}

	fmt.Println(check(ary))
}

func check(ary []int) string {
	win := [][]int{
		{0, 1, 2, 3, 4},
		{5, 6, 7, 8, 9},
		{10, 11, 12, 13, 14},
		{15, 16, 17, 18, 19},
		{20, 21, 22, 23, 24},
		{0, 5, 10, 15, 20},
		{1, 6, 11, 16, 21},
		{2, 7, 12, 17, 22},
		{3, 8, 13, 18, 23},
		{4, 9, 14, 19, 24},
		{0, 6, 12, 18, 24},
		{4, 8, 12, 16, 20},
	}

	for _, w := range win {
		var sum int
		for _, i := range w {
			sum += ary[i]
		}
		switch sum {
		case 5:
			return "O"
		case -5:
			return "X"
		}
	}
	return "D"
}
