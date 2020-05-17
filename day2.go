package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func extract(ln string) (int, int, int) {
	dims := strings.Split(ln, "x")
	w, err := strconv.Atoi(dims[0])
	check(err)
	h, err := strconv.Atoi(dims[1])
	check(err)
	b, err := strconv.Atoi(dims[2])
	check(err)
	ints := []int{w, h, b}
	sort.Ints(ints)
	return ints[0], ints[1], ints[2]
}

func wrapping(ln string) int {
	x, y, z := extract(ln)
	r := 3*x*y + 2*x*z + 2*y*z
	return r
}

func day2() {
	file, err := os.Open("day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	c := 0
	for scanner.Scan() {

		c += wrapping(scanner.Text())

	}
	fmt.Println(c)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
