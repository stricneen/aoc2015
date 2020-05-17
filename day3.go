package main

import "fmt"

// Day3 is here
func Day3() {
	data := input("data/day3.txt")
	set := make(map[pair]bool)

	m := pair{0, 0}
	set[m] = true

	for _, r := range data {
		m = move(r, m)
		set[m] = true
	}

	fmt.Println(len(set))
}

func move(r rune, p pair) pair {
	a := p.a.(int)
	b := p.b.(int)

	if r == '^' {
		a++
	}
	if r == 'v' {
		a--
	}
	if r == '>' {
		b++
	}
	if r == '<' {
		b--
	}
	return pair{a, b}
}
