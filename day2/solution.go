package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type movement struct {
	direction string
	unit      int
}

type position struct {
	horizontal, depth, aim int
}

func (pos *position) updatePosition1(mov movement) {
	switch mov.direction {
	case "forward":
		pos.horizontal += mov.unit
	case "up":
		pos.depth -= mov.unit
	case "down":
		pos.depth += mov.unit
	}
}

func (pos *position) updatePosition2(mov movement) {
	switch mov.direction {
	case "forward":
		pos.horizontal += mov.unit
		pos.depth += (pos.aim * mov.unit)
	case "up":
		pos.aim -= mov.unit
	case "down":
		pos.aim += mov.unit
	}
}

func importFile(file string) []movement {
	raw, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	data := strings.Split(strings.TrimSpace(string(raw)), "\n")
	movements := make([]movement, len(data))
	for idx, el := range data {
		movements[idx] = formatMovement(el)
	}

	return movements
}

func formatMovement(str string) movement {
	strArr := strings.Fields(str)
	num, _ := strconv.Atoi(strArr[1])
	return movement{strArr[0], num}
}

func main() {
	movements := importFile("./input.txt")
	position1 := position{0, 0, 0}
	for _, movement := range movements {
		position1.updatePosition1(movement)
	}
	fmt.Println(position1.horizontal * position1.depth) // 1488669

	position2 := position{0, 0, 0}
	for _, movement := range movements {
		position2.updatePosition2(movement)
	}
	fmt.Println(position2.horizontal * position2.depth) // 1176514794
}
