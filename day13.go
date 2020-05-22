package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type seating struct {
	diner  string
	nextTo string
	happy  int
}

func parse13(ln string) seating {
	s := strings.Split(ln, " ")
	t, _ := strconv.Atoi(s[3])
	if s[2] == "lose" {
		t *= -1
	}
	return seating{s[0], s[10], t}
}

func onScan13(scan *bufio.Scanner) {

	var s []seating

	for scan.Scan() {
		ln := scan.Text()
		if ln == "#" {
			break
		}

		s = append(s, parse13(ln))
	}

	fmt.Println(s)
}

// Day13 is here
func Day13() {
	scan("data/day13.txt", onScan13)
}
