package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func hasThreeVowels(ln string) bool {
	as := strings.Count(ln, "a")
	es := strings.Count(ln, "e")
	is := strings.Count(ln, "i")
	os := strings.Count(ln, "o")
	us := strings.Count(ln, "u")
	return as+es+is+os+us > 2
}

func isValid(ln string) bool {
	as := strings.Count(ln, "ab")
	es := strings.Count(ln, "cd")
	is := strings.Count(ln, "pq")
	us := strings.Count(ln, "xy")
	return as+es+is+us == 0
}

func hasPair(ln string) bool {
	c := ' '
	for _, rune := range ln {
		if rune == c {
			return true
		}
		c = rune
	}
	return false
}

func isNice(ln string) bool {
	return hasThreeVowels(ln) && isValid(ln) && hasPair(ln)
}

func onScan(scan *bufio.Scanner) {
	c := 0
	for scan.Scan() {
		ln := scan.Text()
		nice := isNice(ln)
		if nice {
			c++
		}
	}
	fmt.Println(c)
}

// Day5 is here
func Day5() {
	file, err := os.Open("data/day5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	onScan(scanner)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
