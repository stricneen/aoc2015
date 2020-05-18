package main

import (
	"bufio"
	"fmt"
	"strings"
)

func parse8(ln string) (int, int) {
	var x, y int
	for _, r := range ln {
		x++
		y++
		if r == '\\' {
			y--
		}

	}

	ds := strings.Count(ln, "\\\\")
	xs := strings.Count(ln, "\\x")
	xs2 := strings.Count(ln, "\\\\x")

	return x, y - 2 + ds - (xs * 2) + xs2
}

//1377
//1505
func onScan8(scan *bufio.Scanner) {

	var x, y int

	for scan.Scan() {
		ln := scan.Text()

		xx, yy := parse8(ln)
		x += xx
		y += yy
		fmt.Println(ln)
		fmt.Println(xx, yy)
	}
	fmt.Println(x, y)
	fmt.Println(x - y)
}

// Day8 is here
func Day8() {
	scan("data/day8.txt", onScan8)
}
