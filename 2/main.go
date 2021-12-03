package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(file string) []string {

	data, err := os.Open(file)

	if err != nil {
		panic(err)
	}

	var lines []string
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

func calcFinalPosition(data []string) int {
	horizontal := 0
	depth := 0
	aim := 0

	for _, j := range data {
		line := strings.Fields(j)
		num, err := strconv.Atoi(line[1])

		if err != nil {
			panic(err)
		}

		switch line[0] {
		case "forward":
			horizontal += num
			depth += num * aim

		case "down":
			aim += num

		case "up":
			aim -= num
		}
	}
	return horizontal * depth
}

func main() {
	data := readFile("./day2input.txt")
	fmt.Println(calcFinalPosition(data))
}
