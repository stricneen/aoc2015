package main

import (
	"bufio"
	"fmt"
)

func onScan19(scan *bufio.Scanner) {

	for scan.Scan() {
		ln := scan.Text()
		if ln == "#" {
			break
		}
		fmt.Println(ln)
	}

}

// Day19 is here
func Day19() {
	scan("data/day19.txt", onScan19)
}
