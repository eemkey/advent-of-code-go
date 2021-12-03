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

func getCommonBit(arr []string, isGamma bool) int {
	var result int
	count := []int{0, 0}
	for _, bit := range arr {
		idx, _ := strconv.Atoi(bit)
		count[idx] += 1
	}

	if count[0] == count[1] {
		if isGamma {
			result = 1
		} else {
			result = 0
		}
	} else {
		if isGamma {
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

func findRating(input [][]string, isGamma bool) string {
	result := ""
	for idx := range input[0] {
		currColArr := getColBits(input, idx)
		result += strconv.Itoa(getCommonBit(currColArr, isGamma))
	}

	return result
}

func main() {
	data := importFile("./input.txt")
	gammaRate, _ := strconv.ParseInt(findRating(data, true), 2, 32)
	epsilonRate, _ := strconv.ParseInt(findRating(data, false), 2, 32)

	fmt.Println(gammaRate * epsilonRate) // 4160394
}
