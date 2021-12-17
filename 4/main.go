package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed day4input.txt
var input string

func partOne(input string) int {
	allBoards, bingoResults := parseInput(input)
	finalWinningScore := iterateThroughBoards(allBoards, bingoResults)
	return finalWinningScore

}

func parseInput(input string) ([][][]int, []int) {
	total := strings.Split(input, "\n")

	var bingoResults []int
	bingoTeller := strings.Split(total[0], ",")
	for _, i := range bingoTeller {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		bingoResults = append(bingoResults, j)
	}

	var allBoards [][][]int
	var singleBoard [][]int
	for i, j := range total {
		// In our sample input, the first line is the bingoTeller line, the second line is a blank line.
		if i < 2 {
			continue
		}

		// if we have a blank line, that's the end of a board.
		// add the finalized 5x5 board to allBoards array and reset the single board to a zeroed array.
		if j == "" {
			allBoards = append(allBoards, singleBoard)
			singleBoard = make([][]int, 0)
			continue
		}

		rowStrings := strings.Fields(j)
		row := make([]int, 0)
		for _, number := range rowStrings {
			rowInts, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}
			row = append(row, rowInts)
		}
		singleBoard = append(singleBoard, row)
	}
	allBoards = append(allBoards, singleBoard)
	return allBoards, bingoResults
}

func determineBoardState(board [][]int, bingoResult int) [][]int {
	for a, line := range board {
		for b, char := range line {
			if char == bingoResult {
				board[a][b] = -1
			}
		}
	}
	return board
}

func determineIsWinningRow(board [][]int) bool {
	for i := 0; i < len(board); i++ {
		rowWon := true
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] != -1 {
				rowWon = false
				break
			}
		}
		if rowWon {
			return true
		}
	}
	return false
}

func determineIsWinningColumn(board [][]int) bool {
	for col := 0; col < len(board[0]); col++ {
		colWon := true
		for i := 0; i < len(board); i++ {
			if board[i][col] != -1 {
				colWon = false
				break
			}
		}
		if colWon {
			return true
		}
	}
	return false
}
func iterateThroughBoards(allBoards [][][]int, bingoResults []int) int {
	for _, board := range allBoards {
		for _, bingoResult := range bingoResults {
			board = determineBoardState(board, bingoResult)
			hasWonRow := determineIsWinningRow(board)
			if hasWonRow {
				fmt.Println("WON BY ROW")
				fmt.Println(bingoResult)
				fmt.Println(board)
				return calculateWinningScore(bingoResult, board)
			}
			hasWonColumn := determineIsWinningColumn(board)
			if hasWonColumn {
				fmt.Println("WON BY COL")
				fmt.Println(bingoResult)
				fmt.Println(board)
				return calculateWinningScore(bingoResult, board)
			}
		}
	}
	return 0
}
func calculateWinningScore(bingoResult int, board [][]int) int {
	var score int
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] != -1 {
				score += board[i][j]
			}
		}
	}
	return score * bingoResult
}

func main() {
	fmt.Println("Final Score for part 1: ", partOne(input))
}
