package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)

	sc.Scan()
	s := sc.Text()
	sc.Scan()
	t := sc.Text()
	if s == t {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
