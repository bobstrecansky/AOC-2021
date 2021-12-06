package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed day3input.txt
var input string

func partOne(input string) int {
	var gamma, epsilon string
	binaries := strings.Split(strings.TrimRight(input, "\n"), "\n")
	for i := 0; i < len(binaries[0]); i++ {
		var zeroes, ones int
		for _, j := range binaries {
			if j[i] == '0' {
				zeroes++
			} else {
				ones++
			}
		}

		if zeroes > ones {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	e, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		panic(err)
	}

	g, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		panic(err)
	}

	return int(e * g)
}

func iterateOverBinaries(input string, calcType string) int {
	binaries := strings.Split(strings.TrimRight(input, "\n"), "\n")
	for i := 0; len(binaries) > 1; i++ {
		var zeroes, ones int
		for _, j := range binaries {
			if j[i] == '0' {
				zeroes++
			} else {
				ones++
			}
		}

		var leftoverChar string

		if calcType == "oxygen" {
			leftoverChar = "1"
			if zeroes > ones {
				leftoverChar = "0"
			}
		} else { // Carbon Dioxide
			leftoverChar = "0"
			if zeroes > ones {
				leftoverChar = "1"
			}
		}

		var distilledBinaries []string
		for _, k := range binaries {
			if string(k[i]) == leftoverChar {
				distilledBinaries = append(distilledBinaries, k)
			}
		}
		binaries = distilledBinaries
	}

	result, _ := strconv.ParseInt(binaries[0], 2, 64)
	return int(result)
}

func partTwo(input string) int {
	oxygen := iterateOverBinaries(input, "oxygen")
	carbonDioxide := iterateOverBinaries(input, "carbonDioxide")
	return oxygen * carbonDioxide
}

func main() {
	partOneResult := partOne(input)
	fmt.Println("Part One Output: ", partOneResult)
	partTwoResult := partTwo(input)
	fmt.Println("Part Two Output: ", partTwoResult)
}
