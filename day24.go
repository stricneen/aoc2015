package main

import (
	"bufio"
	"fmt"
	"strconv"
)

func exec24(p []int) {
	t := 0
	for _, v := range p {
		t += v
	}

	fmt.Println("Total", t, "Group", t/3)

}

func onScan24(scan *bufio.Scanner) {

	p := make([]int, 0)
	for scan.Scan() {
		ln := scan.Text()
		t, _ := strconv.Atoi(ln)
		p = append(p, t)

		if ln == "#" {
			break
		}
		//fmt.Println(ln)
	}
	exec24(p)

}

// Day24 is here
func Day24() {
	scan("data/day24.txt", onScan24)
}
