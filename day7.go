package main

import (
	"bufio"
	"fmt"
)

type wire struct {
	name  string
	value int
}

func onScan7(scan *bufio.Scanner) {

	for scan.Scan() {
		ln := scan.Text()
		if ln == "#" {
			break
		}
		fmt.Println(ln)
	}
}

// Day7 is here
func Day7() {
	scan("data/day7.txt", onScan7)
}
