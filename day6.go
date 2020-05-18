package main

import (
	"bufio"
	"fmt"
)

func onScan6(scan *bufio.Scanner) {

	var lights [1000][1000]bool

	for scan.Scan() {
		ln := scan.Text()
		fmt.Println(ln)
	}

	var c, i, j int
	for i = 0; i < 1000; i++ {
		for j = 0; j < 1000; j++ {
			if lights[i][j] {
				c++
			}
		}
	}
	fmt.Println(c)
}

// Day6 is here
func Day6() {
	scan("data/day6.txt", onScan6)
}
