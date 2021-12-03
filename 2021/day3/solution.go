package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func importFile(file string) [][]string {
	raw, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	data := strings.Split(strings.TrimSpace(string(raw)), "\n")
	numbers := make([][]string, len(data))
	for idx, bits := range data {
		numbers[idx] = strings.Split(bits, "")
	}
	return numbers
}

func getColBits(arr [][]string, idx int) []string {
	var result []string
	for _, num := range arr {
		result = append(result, num[idx])
	}
	return result
}

func getCommonBit(arr []string, isMost bool) int {
	var result int
	count := []int{0, 0}
	for _, bit := range arr {
		idx, _ := strconv.Atoi(bit)
		count[idx] += 1
	}

	if count[0] == count[1] {
		if isMost {
			result = 1
		} else {
			result = 0
		}
	} else {
		if isMost {
			if count[0] > count[1] {
				result = 0
			} else {
				result = 1
			}
		} else {
			if count[0] > count[1] {
				result = 1
			} else {
				result = 0
			}
		}
	}
	return result
}

func findRating1(input [][]string, isMost bool) string {
	result := ""
	for idx := range input[0] {
		currColArr := getColBits(input, idx)
		result += strconv.Itoa(getCommonBit(currColArr, isMost))
	}

	return result
}

func filterNums(arr [][]string, idx, bit int) [][]string {
	var result [][]string
	for _, num := range arr {
		if num[idx] == strconv.Itoa(bit) {
			result = append(result, num)
		}
	}
	return result
}

func findRating2(input [][]string, isMost bool) string {
	arr := input[:]
	idx := 0

	for len(arr) > 1 {
		colBits := getColBits(arr, idx)
		mostCommonBit := getCommonBit(colBits, isMost)
		arr = filterNums(arr, idx, mostCommonBit)

		idx += 1
	}

	return strings.Join(arr[0], "")
}

func main() {
	data := importFile("./input.txt")
	gammaRate, _ := strconv.ParseInt(findRating1(data, true), 2, 32)
	epsilonRate, _ := strconv.ParseInt(findRating1(data, false), 2, 32)

	fmt.Println(gammaRate * epsilonRate) // 4160394

	oxygenRating, _ := strconv.ParseInt(findRating2(data, true), 2, 32)
	co2Rating, _ := strconv.ParseInt(findRating2(data, false), 2, 32)

	fmt.Println(oxygenRating * co2Rating) // 4125600
}
