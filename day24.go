package main

import (
	"bufio"
	"fmt"
	"strconv"
)

// type pair struct {
// 	x int
// 	y int
// }

func pairs(sum int, set []int) [][]int {
	r := make([][]int, 0)
	for i := 0; i < len(set); i++ {
		for j := i + 1; j < len(set); j++ {
			if set[i]+set[j] == sum {
				r = append(r, []int{set[i], set[j]})
			}

		}
	}

	return r
}

func exec24(p []int) {
	t := 0
	for _, v := range p {
		t += v
	}

	fmt.Println("Total", t, "Group", t/3)

	ps := pairs(t/3, p)
	fmt.Println(ps)

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
