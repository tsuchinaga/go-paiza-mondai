package main

import (
	"fmt"
)

func main() {
	var n, m int
	var s string
	_, _ = fmt.Scan(&n, &m)

	dots := make([][]int, n)
	for i := range dots {
		dots[i] = make([]int, m)

		_, _ = fmt.Scan(&s)
		for j, r := range []rune(s) {
			if r == '#' {
				dots[i][j] = 1
			}
		}
	}

	ans := 0
	// 線対象
	if isLine(n, m, dots) {
		ans += 1
	}

	// 点対象
	// 180度回転させたときに同じになる
	// 後ろから真ん中まで順にみていって、すべて一致していれば点対象
	if isPoint(n, m, dots) {
		ans += 2
	}

	switch ans {
	case 0:
		fmt.Println("none")
	case 1:
		fmt.Println("line symmetry")
	case 2:
		fmt.Println("point symmetry")
	case 3:
		fmt.Println("line point symmetry")
	}
}

func isLine(n, m int, dots [][]int) bool {
	symmetry := true
	// 縦の線対象
	// 縦線に対して左右が一致している
	// 左右から順に真ん中までみていって、全行で一致していれば線対象
	for i := 0; i < n; i++ {
		for j := 0; j < m/2; j++ {
			if dots[i][j] != dots[i][m-1-j] {
				symmetry = false
				break
			}
		}
		if !symmetry {
			break
		}
	}
	if symmetry {
		return true
	}

	symmetry = true
	// 横の線対処
	// 横線に対して上下が一致している
	// 上下から順に真ん中までみていって、全列で一致していれば線対象
	// 縦の線対象の場合はチェックをスキップできる
	for i := 0; i < m; i++ {
		for j := 0; j < n/2; j++ {
			if dots[j][i] != dots[n-1-j][i] {
				symmetry = false
				break
			}
		}
		if !symmetry {
			break
		}
	}
	return symmetry
}

func isPoint(n, m int, dots [][]int) bool {
	for k := 0; k < (n*m)/2; k++ {
		i, j := k/m, k%m
		if dots[i][j] != dots[(n-1)-i][(m-1)-j] {
			return false
		}
	}
	return true
}
