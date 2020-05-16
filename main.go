package main

import (
	//   "io"
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	day2()

	//fmt.Println(required("2x3x4"))
	//fmt.Println(required("1x1x10"))
}

func required(ln string) int {
	//fmt.Println(ln)
	dims := strings.Split(ln, "x")
	w, err := strconv.Atoi(dims[0])
	check(err)
	h, err := strconv.Atoi(dims[1])
	check(err)
	b, err := strconv.Atoi(dims[2])
	check(err)

	x := w * h
	y := b * h
	z := w * b

	extra := min(x, y, z)

	r := (2 * x) + (2 * y) + (2 * z) + extra
	//fmt.Println(r)

	return r
}

func min(a, b, c int) int {
	if a < b && a < c {
		return a
	}
	if b < a && b < c {
		return b
	}
	return c
}

// 1608378

func day2() {
	file, err := os.Open("day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	c := 0
	for scanner.Scan() {

		r := required(scanner.Text())
		c += r
		// fmt.Println(scanner.Text())
	}
	fmt.Println(c)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func day1() {
	data := input()

	// part 1
	up := strings.Count(data, "(")
	down := strings.Count(data, ")")
	fmt.Println(up - down)

	// part 2
	c := 0
	for i, rune := range data {
		if c < 0 {
			fmt.Printf("%d\n", i)
			break
		}
		if rune == '(' {
			c++
		}
		if rune == ')' {
			c--
		}
	}
}

func input() string {
	dat, err := ioutil.ReadFile("day1.txt")
	check(err)
	return string(dat)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
