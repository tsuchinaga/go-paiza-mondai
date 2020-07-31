package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	bs := bufio.NewScanner(os.Stdin)
	bs.Split(bufio.ScanLines)
	bs.Buffer(make([]byte, 100000*4), 100000*4)
	bs.Scan()
	s := bs.Text()
	rs := []rune(s)
	var cnt int
	for i := 0; i < len(rs)-1; i++ {
		if string(rs[i])+string(rs[i+1]) == "令和" {
			cnt++
		}
	}
	fmt.Println(cnt)
}
