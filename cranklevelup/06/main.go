package main

import (
	"fmt"
)

func main() {
	var h int
	_, _ = fmt.Scan(&h)

	dah := make([]int, 2) // 過去2回の自分 -> 敵へのダメージ
	dbh := make([]int, 2) // 過去2回の敵 -> 自分へのダメージ
	n := 0                // nターン目
	for h > 0 {
		n++

		// 自分から敵への攻撃
		da := 1
		if n > 2 {
			da = dbh[0] + dbh[1]
		}

		// 敵から自分への攻撃
		db := 1
		if n > 2 {
			db = dah[0] + dah[1] + dah[1]
		}

		// 状況更新
		dah[0], dah[1] = dah[1], da
		dbh[0], dbh[1] = dbh[1], db
		h -= db
	}

	fmt.Println(n)
}
