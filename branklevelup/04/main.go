package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bs := bufio.NewScanner(os.Stdin)
	bs.Split(bufio.ScanLines)

	bs.Scan()
	l1s := strings.Split(bs.Text(), " ")
	a, _ := strconv.Atoi(l1s[0])
	b, _ := strconv.Atoi(l1s[1])

	bs.Scan()
	for i, c := range []rune(bs.Text()) {
		cs := string(c)
		if cs != "" && a-1 <= i && i < b {
			fmt.Print(strings.ToUpper(cs))
		} else {
			fmt.Print(cs)
		}
	}
}
