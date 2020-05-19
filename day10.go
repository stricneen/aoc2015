package main

import (
	"fmt"
	"strconv"
)

// Day10 is here
func Day10() {

	start := "3113322113"
	r := ""
	ps(start)

	for i := 0; i < 40; i++ {
		r = step(start)
		//ps(r)
		pi(i)
		start = r
	}

	fmt.Println(len(start))
	//end := step(start)
	// ps(end)
}

func step(in string) string {
	ret := ""
	var cur rune = ' '
	c := 0

	for _, r := range in {
		if r != cur && cur != ' ' {
			ret += strconv.Itoa(c) + string(cur)
			c = 0
		}
		cur = r
		c++
	}
	ret += strconv.Itoa(c) + string(cur)
	return ret
}
