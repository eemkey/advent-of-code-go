package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type board [5][5]square

type square struct {
	num    int
	marked bool
}

func importFile(file string) ([]int, []board) {
	raw, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	data := strings.Split(strings.TrimSpace(string(raw)), "\n\n")
	numsStr := data[0]
	boardsSlice := data[1:]
	drawnNums := convertNumsStr(numsStr)
	boards := convertBoardsStr(boardsSlice, len(data)-1)

	return drawnNums, boards
}

func convertNumsStr(str string) []int {
	sliceStr := strings.Split(str, ",")
	sliceInt := make([]int, len(sliceStr))
	for idx, str := range sliceStr {
		num, _ := strconv.Atoi(str)
		sliceInt[idx] = num
	}
	return sliceInt
}

func convertBoardsStr(boardsSlice []string, numOfBoards int) []board {
	boards := make([]board, numOfBoards)
	for idx := 0; idx < len(boards); idx++ {
		boards[idx].populateBoard(boardsSlice[idx])
	}

	return boards
}

func (b *board) populateBoard(boardStr string) {
	boardSlice := strings.Split(boardStr, "\n")
	for rowIdx, row := range boardSlice {
		for colIdx, num := range strings.Fields(row) {
			b[rowIdx][colIdx].num, _ = strconv.Atoi(num)
		}
	}
}

func getSumOfUnmarked(b board) int {
	sum := 0
	for _, row := range b {
		for _, col := range row {
			if !col.marked {
				sum += col.num
			}
		}
	}
	return sum
}

func drawNums1(nums []int, boards []board) int {
	for _, drawnNum := range nums {
		for boardsIdx := 0; boardsIdx < len(boards); boardsIdx++ {
			isMarked := boards[boardsIdx].markBoard(drawnNum)
			if isMarked {
				if boards[boardsIdx].isWinner() {
					sum := getSumOfUnmarked(boards[boardsIdx])
					return sum * drawnNum
				}
			}
		}
	}
	return 0
}

func drawNums2(nums []int, boards []board) int {
	winners := make([]bool, len(boards))
	for _, drawnNum := range nums {
		for boardsIdx := 0; boardsIdx < len(boards); boardsIdx++ {
			isMarked := boards[boardsIdx].markBoard(drawnNum)
			if isMarked {
				if boards[boardsIdx].isWinner() {
					winners[boardsIdx] = true
					if allWinners(winners) {
						sum := getSumOfUnmarked(boards[boardsIdx])
						return sum * drawnNum
					}
				}
			}
		}
	}
	return 0
}

func allWinners(s []bool) bool {
	count := 0
	for _, el := range s {
		if el {
			count += 1
		}
	}
	return count == len(s)
}

func (b *board) markBoard(drawnNum int) bool {
	for rowIdx := 0; rowIdx < len(b); rowIdx++ {
		for colIdx := 0; colIdx < len(b[rowIdx]); colIdx++ {
			if b[rowIdx][colIdx].num == drawnNum {
				b[rowIdx][colIdx].marked = true
				return true
			}
		}
	}
	return false
}

func (b board) isWinner() bool {
	if checkRows(b) || checkCols(b) {
		return true
	}

	return false
}

func checkCols(b board) bool {
	for rowIdx, row := range b {
		colSlice := make([]square, len(row))
		for colIdx := range row {
			colSlice[colIdx] = b[colIdx][rowIdx]
		}
		if checkLine(colSlice) {
			return true
		}
	}
	return false
}

func checkRows(b board) bool {
	for _, row := range b {
		rowSlice := row[:]
		if checkLine(rowSlice) {
			return true
		}
	}
	return false
}

func checkLine(line []square) bool {
	for _, square := range line {
		if !square.marked {
			return false
		}
	}
	return true
}

func main() {
	drawnNums, boards := importFile("./input.txt")
	firstPlace := drawNums1(drawnNums, boards)
	lastPlace := drawNums2(drawnNums, boards)
	fmt.Println(firstPlace, lastPlace)
}
