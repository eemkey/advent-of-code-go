package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func importFile(file string) []int {
	raw, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	data := strings.Split(strings.TrimSpace(string(raw)), "\n")
	depths := make([]int, len(data))

	for idx, el := range data {
		num, _ := strconv.Atoi(el)
		depths[idx] = num
	}
	return depths
}

func getSumOfWindow(window []int) int {
	sum := 0
	for _, num := range window {
		sum += num
	}
	return sum
}

func measurements1(depths []int) int {
	count := 0
	for idx := 1; idx < len(depths); idx++ {
		if depths[idx] > depths[idx-1] {
			count += 1
		}
	}
	return count
}

func measurements2(depth []int) int {
	count := 0
	var firstWindow int
	var secondWindow int
	for idx := 0; idx < len(depth)-3; idx++ {
		firstWindow = getSumOfWindow(depth[idx : idx+3])
		secondWindow = getSumOfWindow((depth[idx+1 : idx+4]))
		if secondWindow > firstWindow {
			count += 1
		}
	}
	return count
}

func main() {
	depths := importFile("./input.txt")
	count1 := measurements1(depths)
	fmt.Println(count1) // 1475

	count2 := measurements2(depths)
	fmt.Println(count2) // 1516
}
