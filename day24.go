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

func pairs(group int, sum int, set []int) [][]int {
	r := make([][]int, 0)
	for i := 0; i < len(set); i++ {
		for j := i + 1; j < len(set); j++ {
			if set[i]+set[j] == sum {
				r = append(r, []int{group - sum, set[i], set[j]})
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
	group := t / 3

	fmt.Println("Total", t, "Group", group)

	//	ps := pairs(t/3, p)

	for i := 0; i < len(p); i++ {

		ps := pairs(group, group-p[i], remove(p[i], p))

		//fmt.Println(group - p[i])
		fmt.Println(ps)
		fmt.Println()
	}

}

func remove(i int, is []int) []int {
	c := 0
	r := make([]int, len(is)-1)
	for _, v := range is {
		if v != i {
			r[c] = v
			c++
		}
	}
	return r
}

func onScan24(scan *bufio.Scanner) {

	p := make([]int, 0)
	for scan.Scan() {
		ln := scan.Text()
		if ln == "#" {
			break
		}

		t, _ := strconv.Atoi(ln)
		p = append(p, t)
		//fmt.Println(ln)
	}
	exec24(p)

}

// Day24 is here
func Day24() {
	scan("data/day24.txt", onScan24)
}
