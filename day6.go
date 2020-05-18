package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func onScan6(scan *bufio.Scanner) {

	var lights [1000][1000]bool

	for scan.Scan() {
		ln := scan.Text()
		cmd, x1, y1, x2, y2 := parse(ln)

		fmt.Println(ln)
		fmt.Println(cmd, x1, y1, x2, y2)
	}

	howMany(lights)

}

func parseCoords(coords string) (int, int) {
	c := strings.Split(coords, ",")
	x, err := strconv.Atoi(c[0])
	check(err)
	y, err := strconv.Atoi(c[1])
	check(err)
	return x, y
}

func parse(ln string) (string, int, int, int, int) {
	s := strings.Split(ln, " ")

	if s[0] == "toggle" {
		x1, y1 := parseCoords(s[1])
		x2, y2 := parseCoords(s[3])
		return "toggle", x1, y1, x2, y2
	} else {
		x1, y1 := parseCoords(s[2])
		x2, y2 := parseCoords(s[4])
		return s[1], x1, y1, x2, y2
	}

}

func howMany(mat [1000][1000]bool) {
	var c, i, j int
	for i = 0; i < 1000; i++ {
		for j = 0; j < 1000; j++ {
			if mat[i][j] {
				c++
			}
		}
	}
	fmt.Println(c)
}

// Day6 is here
func Day6() {
	scan("data/day6.txt", onScan6)
}
