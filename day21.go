package main

import "fmt"

type player struct {
	name      string
	hitPoints int
	damage    int
	armor     int
}

// Day21 is
func Day21() {

	boss := player{"boss", 100, 8, 2}
	henry := player{"henry", 100, 0, 0}

	winner := fight(boss, henry)

	fmt.Println(winner)
}

func fight(p1 player, p2 player) player {

	return p1
}
