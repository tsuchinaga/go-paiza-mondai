package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, a int
	_, _ = fmt.Scan(&n)

	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&a)
		if a%2 == 1 {
			nums = append(nums, a)
		}
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for _, a := range nums {
		fmt.Println(a)
	}
}
