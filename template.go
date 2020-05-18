package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// )

// func onScan(scan *bufio.Scanner) {

// 	for scan.Scan() {
// 		ln := scan.Text()
// 		fmt.Println(ln)
// 	}

// }

// // Dayx is here
// func Dayx() {
// 	file, err := os.Open("data/day5.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	onScan(scanner)

// 	if err := scanner.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// }
