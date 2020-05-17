package main

import (
	"io/ioutil"
)

type pair struct {
	a, b interface{}
}

func input(filename string) string {
	dat, err := ioutil.ReadFile(filename)
	check(err)
	return string(dat)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}