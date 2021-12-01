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

	// This secondary depth value will be used as for our first of two comparisons for each row scan.
	scanner.Scan()
	secondaryDepth, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(fmt.Errorf("could not convert secondaryDepth"))
	}

	if secondaryDepth > initialDepth {
		numIncreases++
	}

	for scanner.Scan() {
		// This tertiary depth value will be used as the second of two comparisons for each row scan.
		tertiaryDepth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(fmt.Errorf("could not convert tertiaryDepth"))
		}

		if tertiaryDepth > secondaryDepth {
			numIncreases++
		}
		secondaryDepth = tertiaryDepth
	}

	fmt.Println(numIncreases)
}
