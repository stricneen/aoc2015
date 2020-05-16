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
}

func required(ln string) int {
	fmt.Println(ln)
	x := strings.Split(ln, "x")
	i, err := strconv.Atoi(x[0])
	check(err)
	return i
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
