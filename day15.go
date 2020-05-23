package main

import "fmt"

// Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
// Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3

func score(in []ingredient) int {

	s := ingredient{"score", 0, 0, 0, 0, 0, 0}

	for _, i := range in {
		s.capacity += i.capacity * i.amount
		s.durability += i.durability * i.amount
		s.flavor += i.flavor * i.amount
		s.texture += i.texture * i.amount
	}

	s.capacity = pos(s.capacity)
	s.durability = pos(s.durability)
	s.flavor = pos(s.flavor)
	s.texture = pos(s.texture)

	return s.capacity * s.durability * s.flavor * s.texture
}

// 62842880

func pos(i int) int {
	if i < 0 {
		return 0
	}
	return i
}

// Day15 is
func Day15() {

	test15[0].amount = 44
	test15[1].amount = 56

	s := score(test15)
	fmt.Println(s)

}

var test15 = []ingredient{
	ingredient{"Butterscotch", -1, -2, 6, 3, 8, 0},
	ingredient{"Cinnamon", 2, 3, -2, -1, 3, 0},
}

var ingredients = []ingredient{
	ingredient{"Sprinkles", 5, -1, 0, 0, 5, 0},
	ingredient{"PeanutButter", -1, 3, 0, 0, 1, 0},
	ingredient{"Frosting", 0, -1, 4, 0, 6, 0},
	ingredient{"Sugar", -1, 0, 0, 2, 8, 0},
}

type ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
	amount     int
}
