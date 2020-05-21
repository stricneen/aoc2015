package main

import (
	"fmt"
	"strings"
)

// Day11 is here
func Day11() {
	start := "vzbxkghb"

	ps(start)
	n := start
	for i := 0; i < 30; i++ {
		n = next(n)
		ps(n)
	}
}

func isValidPassword(in string) bool {
	return strings.Contains(in, "q")
}

func next(in string) string {
	n := in

	valid := false
	for valid == false {
		n = inc(n)
		n = ilo(n)
		valid = isValidPassword(n)
		fmt.Println(valid, n)
	}
	return n
}

func ilo(in string) string {
	ret := ""
	f := false
	for k := range in {
		x := in[k]
		if f {
			x = 97
		} else if x == 105 || x == 108 || x == 111 {
			x++
			f = true
		}
		ret += string(x)
	}
	return ret
}

func inc(in string) string {
	ret := ""
	carry := true
	for k := range in {
		x := in[len(in)-1-k]
		n := x
		if carry {
			n++
			if n == 123 {
				n = 97
			} else {
				carry = false
			}
		}
		ret = string(n) + ret
	}
	return ret
}
