package main

import "fmt"

// Day3 is here
func Day3() {
	data := input("data/day3.txt")
	set := make(map[pair]bool)

	a := 0
	b := 0
	m := pair{a, b}
	set[m] = true

	for _, rune := range data {
		//fmt.Println(rune)
		if rune == '^' {
			a++
		}
		if rune == '>' {
			b++
		}
		if rune == '<' {
			b--
		}
		if rune == 'v' {
			a--
		}

		n := pair{a, b}
		set[n] = true
	}

	fmt.Println(len(set))

}

// func move(rune int, pos pair) pair {
// 	a := pos.a
// 	b := pos.b

// 	if rune == '^' {
// 		return pair{a + 1, b}
// 	}

// 	return pos
// }
