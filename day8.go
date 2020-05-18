package main

import (
	"bufio"
	"fmt"
)

func onScan8(scan *bufio.Scanner) {

	for scan.Scan() {
		ln := scan.Text()
		fmt.Println(ln)
	}

}

// Day8 is here
func Day8() {
	scan("data/day8.txt", onScan8)
}
