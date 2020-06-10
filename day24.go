package main

import (
	"bufio"
	"fmt"
)

func onScan24(scan *bufio.Scanner) {

	for scan.Scan() {
		ln := scan.Text()
		if ln == "#" {
			break
		}
		fmt.Println(ln)
	}

}

// Day24 is here
func Day24() {
	scan("data/day24.txt", onScan24)
}
