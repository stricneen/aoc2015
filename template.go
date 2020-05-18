package main

import (
	"bufio"
	"fmt"
)

func onScanT(scan *bufio.Scanner) {

	for scan.Scan() {
		ln := scan.Text()
		fmt.Println(ln)
	}

}

// DayT is here
func DayT() {
	scan("data/day6.txt", onScan6)
}
