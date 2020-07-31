package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	ks := make([]string, n)

	var k string
	for i := range ks {
		_, _ = fmt.Scan(&k)
		ks[i] = k
	}

	var cnt int
	var result bool
	for _, a := range ks {
		for _, b := range ks {
			s := a + b
			if cnt < 1000 {
				fmt.Println(s)
				cnt++
			}
			if s == "令和" {
				result = true
			}
		}
	}

	if result {
		fmt.Println("Nice")
	} else {
		fmt.Println("Bad")
	}
}
