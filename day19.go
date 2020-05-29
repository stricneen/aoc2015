package main

import (
	"bufio"
	"fmt"
	"strings"
)

func overCount(s string, sub string) int {
	c := 0
	for i := 0; i < len(s)-len(sub)+1; i++ {
		ssub := s[i : i+len(sub)]
		if ssub == sub {
			c++
		}
	}
	return c
}

func onScan19(scan *bufio.Scanner) {

	muts := make(map[string]bool)

	start := "ORnPBPMgArCaCaCaSiThCaCaSiThCaCaPBSiRnFArRnFArCaCaSiThCaCaSiThCaCaCaCaCaCaSiRnFYFArSiRnMgArCaSiRnPTiTiBFYPBFArSiRnCaSiRnTiRnFArSiAlArPTiBPTiRnCaSiAlArCaPTiTiBPMgYFArPTiRnFArSiRnCaCaFArRnCaFArCaSiRnSiRnMgArFYCaSiRnMgArCaCaSiThPRnFArPBCaSiRnMgArCaCaSiThCaSiRnTiMgArFArSiThSiThCaCaSiRnMgArCaCaSiRnFArTiBPTiRnCaSiAlArCaPTiRnFArPBPBCaCaSiThCaPBSiThPRnFArSiThCaSiThCaSiThCaPTiBSiRnFYFArCaCaPRnFArPBCaCaPBSiRnTiRnFArCaPRnFArSiRnCaCaCaSiThCaRnCaFArYCaSiRnFArBCaCaCaSiThFArPBFArCaSiRnFArRnCaCaCaFArSiRnFArTiRnPMgArF"

	for scan.Scan() {

		ln := scan.Text()
		if ln == "#" {
			break
		}

		x := strings.Split(ln, " => ")
		//c := strings.Count(start, x[0])

		for i := 0; i < len(start)-len(x[0])+1; i++ {
			ssub := start[i : i+len(x[0])]
			if ssub == x[0] {

				nm := start[0:i] + x[1] + start[i+len(x[0]):]

				muts[nm] = true
			}
		}

	}

	fmt.Println(len(muts))
}

// Day19 is here
func Day19() {
	scan("data/day19.txt", onScan19)
}
