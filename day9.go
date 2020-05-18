package main

import (
	"bufio"
	"fmt"
)

func onScan9(scan *bufio.Scanner) {

	for scan.Scan() {
		ln := scan.Text()
		fmt.Println(ln)
	}

}

// Day9 is here
func Day9() {
	scan("data/day9.txt", onScan9)
}
