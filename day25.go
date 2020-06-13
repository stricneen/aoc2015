package main

import "fmt"

//To continue, please consult the code grid in the manual.
//Enter the code at row 2981, column 3075.

func day25() {
	fmt.Println("Happy Christmas")

	// row := 2981
	// col := 3075

	fmt.Println(number(4, 2))
	fmt.Println(number(1, 5))

}

func number(r int, c int) int {

	ctop := addDown(c)
	acc := c
	for i := 1; i < r; i++ {
		ctop += acc
		acc++
	}

	return ctop
}

func addDown(i int) int {
	if i == 1 {
		return 1
	}
	return i + addDown(i-1)
}
