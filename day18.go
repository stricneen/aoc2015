package main

import (
	"bufio"
	"fmt"
)

func surroundGrid(grid [][]string) [][]string {
	r := len(grid)
	ret := makeGrid(r + 2)
	//	c := len(i[0])
	for x, i := range grid {
		for y := range i {
			ret[x+1][y+1] = grid[x][y]
		}
	}

	return ret
}

func stepGrid(grid [][]string) [][]string {

	// r := len(grid)
	// c := len(grid[0])
	//	ret := makeGrid(len(grid))

	ret := surroundGrid(grid)

	// for x, r := range grid {
	// 	for y := range r {

	// 		//	c := 0
	// 		p := grid[x][y]

	// 		ret[x][y] = p
	// 		// if p != "#" {
	// 		// 	ret[x][y] = "#"
	// 		// } else {
	// 		// 	ret[x][y] = "."
	// 		// }

	// 		//fmt.Println(x, y)
	// 	}

	// }

	return ret
}

func count(grid [][]string) int {
	c := 0
	for _, i := range grid {
		for _, ii := range i {
			if ii == "#" {
				c++
			}
		}
	}
	return c
}

func makeGrid(i int) [][]string {
	a := make([][]string, i)
	for c := range a {
		a[c] = make([]string, i)
	}
	return a
}

func printGrid(i [][]string) {
	fmt.Println()
	for _, r := range i {
		for _, s := range r {
			if len(s) == 1 {

				fmt.Print(s)
			} else {
				fmt.Print("_")
			}
		}
		fmt.Println()
	}
}

func onScan18(scan *bufio.Scanner) {

	grid := makeGrid(6)

	r := 0
	for scan.Scan() {
		ln := scan.Text()
		if ln == "@" {
			break
		}

		for i, x := range ln {
			grid[r][i] = string(x)
		}
		r++
		//	fmt.Println(ln)
	}

	printGrid(grid)
	s1 := stepGrid(grid)
	printGrid(s1)

	c := count(grid)
	fmt.Println(c)
}

// Day18 is here
func Day18() {
	scan("data/day18.txt", onScan18)
}
