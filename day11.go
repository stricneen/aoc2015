package main

import "fmt"

// Day11 is here
func Day11() {

	start := "vzbxkghb"

	n := next(start)

	ps(start)
	ps(n)

}

func next(in string) string {
	ret := ""
	for k := range in {
		x := in[len(in)-1-k]
		if x < 122 {
			ret += string(x + 1)
		} else {
			ret += "a"

		}
		fmt.Println(x)
		// now k starts from the end
	}
	return ret

}
