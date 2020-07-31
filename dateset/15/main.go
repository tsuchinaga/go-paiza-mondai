package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var s string
	_, _ = fmt.Scan(&s)

	// /区切りで3つかどうか
	ymd := strings.Split(s, "/")
	if len(ymd) != 3 {
		fmt.Println("No")
		return
	}

	// 各項目の長さが条件を満たしているか
	if len(ymd[0]) != 4 || len(ymd[1]) != 2 || len(ymd[2]) != 2 {
		fmt.Println("No")
		return
	}

	// y, m, dがそれぞれ数値で条件を満たしているか
	if _, err := strconv.Atoi(ymd[0]); err != nil {
		fmt.Println("No")
		return
	}

	if m, err := strconv.Atoi(ymd[1]); err != nil {
		fmt.Println("No")
		return
	} else if m < 1 || 12 < m {
		fmt.Println("No")
		return
	}

	if d, err := strconv.Atoi(ymd[2]); err != nil {
		fmt.Println("No")
		return
	} else if d < 1 || 31 < d {
		fmt.Println("No")
		return
	}

	fmt.Println("Yes")
}
