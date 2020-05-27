package main

import (
	"bufio"
	"fmt"
)

func onScan18(scan *bufio.Scanner) {

	for scan.Scan() {
		ln := scan.Text()
		if ln == "#" {
			break
		}
		fmt.Println(ln)
	}

}

// Day18 is here
func Day18() {
	scan("data/day18.txt", onScan18)
}
