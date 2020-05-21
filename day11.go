package main

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

func next(in string) string {
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
