package main

import "io/ioutil"

func input() string {
	dat, err := ioutil.ReadFile("data/day1.txt")
	check(err)
	return string(dat)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
