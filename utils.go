package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
)

type pair struct {
	a, b interface{}
}

func input(filename string) string {
	dat, err := ioutil.ReadFile(filename)
	check(err)
	return string(dat)
}

type onScanFunc func(*bufio.Scanner)

func scan(filename string, f onScanFunc) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	f(scanner)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
