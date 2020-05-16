package main

import (
	//   "io"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data := input()

	// part 1
	up := strings.Count(data, "(")
	down := strings.Count(data, ")")
	fmt.Println(up - down)

	// part 2
	c := 0
	for i, rune := range data {
		if c < 0 {
			fmt.Printf("%d\n", i)
			break
		}
		if rune == '(' {
			c += 1
		}
		if rune == ')' {
			c -= 1
		}
	}
}

func input() string {
	dat, err := ioutil.ReadFile("day1.txt")
	check(err)
	return string(dat)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
