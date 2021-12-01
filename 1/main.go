package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var numIncreases int

	file, err := os.Open("day1input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// We are capturing an initial depth value for comparison here
	scanner.Scan()
	initialDepth, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(fmt.Errorf("could not convert initialDepth"))
	}

	scanner.Scan()
	secondaryDepth, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(fmt.Errorf("could not convert secondaryDepth"))
	}

	scanner.Scan()
	tertiaryDepth, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(fmt.Errorf("could not convert tertiaryDepth"))
	}

	for scanner.Scan() {
		quaternaryDepth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(fmt.Errorf("could not convert tertiaryDepth"))
		}

		if initialDepth + secondaryDepth + tertiaryDepth < secondaryDepth + tertiaryDepth + quaternaryDepth {
			numIncreases++
		}

		// Shift each value in the sliding window one to the right
		initialDepth = secondaryDepth
		secondaryDepth = tertiaryDepth
		tertiaryDepth = quaternaryDepth
	}
	fmt.Println(numIncreases)
}
