package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type dist struct {
	from     string
	to       string
	distance int
}

type route struct {
	visit    []string
	distance int
}

// Permute the values at index i to len(a)-1.
func perm(a []string, f func([]string), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

func onScan9(scan *bufio.Scanner) {
	x := make(map[dist]int)
	y := make(map[string]bool)

	for scan.Scan() {
		ln := scan.Text()
		if ln == "#" {
			break
		}
		d := parse9(ln)
		x[d] = d.distance

		y[d.from] = true
		y[d.to] = true
	}

	locations := make([]string, len(y))
	c := 0
	for k := range y {
		locations[c] = k
		c++
	}

	for i := 0; i < len(locations)/2; i++ {
		perm(locations, func(p []string) {

			ft := getDist(p, x)

			fmt.Println(p)
			fmt.Println(ft)

		}, i)
	}
}

func getDist(locs []string, d map[dist]int) int {
	return 0
}

func parse9(ln string) dist {
	p := strings.Split(ln, " ")
	i, _ := strconv.Atoi(p[4])
	return dist{p[0], p[2], i}
}

// Day9 is here
func Day9() {
	scan("data/day9.txt", onScan9)
}
