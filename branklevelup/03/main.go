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
	s := bs.Text()

	fmt.Println(string([]rune(s)[a-1 : b]))
}
