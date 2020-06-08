package main

import (
	"fmt"
)

func main() {
	var p, q, r int // グループAの人数, グループBの人数, グループCの人数
	var i, j int    // i番目の人がj番目の一人に仕事を渡す
	_, _ = fmt.Scan(&p, &q, &r)
	abJob := make([]int, p)
	bcJob := make([]int, q)

	for k := 0; k < p; k++ {
		_, _ = fmt.Scan(&i, &j)
		abJob[i-1] = j - 1
	}

	for k := 0; k < q; k++ {
		_, _ = fmt.Scan(&i, &j)
		bcJob[i-1] = j - 1
	}

	for i := 0; i < p; i++ {
		fmt.Printf("%d %d\n", i+1, bcJob[abJob[i]]+1)
	}
}
