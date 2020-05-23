package main

import (
	"bufio"
	"fmt"
)

func onScan17(scan *bufio.Scanner) {

	for scan.Scan() {
		ln := scan.Text()
		if ln == "#" {
			break
		}
		fmt.Println(ln)
	}

}

// Day17 is here
func Day17() {
	scan("data/day17.txt", onScan17)
}
