package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile(file string) []string {

	var sumOfColumns []int
	var numLines int

	data, err := os.Open(file)

	if err != nil {
		panic(err)
	}

	var lines []string
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		for i, j := range scanner.Text(){
			fmt.Print("I: ", i, " J: ", j-48, "\n")
//			sumOfColumns[i] += in t(j-48) //ascii to decimal
		}
		numLines++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(sumOfColumns)
	fmt.Println(numLines)
	return lines
}

func main() {

data := readFile("./day3sample.txt")
fmt.Println(data)
}
