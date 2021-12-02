package main

import (
	"advent-2021/utils"
	"fmt"
	"strconv"
	"strings"
)

type positions struct {
	horizontal int
	depth      int
	aim        int
}

func main() {
	arr := utils.SetupDay2()
	solvePart1(arr)
	solvePart2(arr)
}
func splitStep(step string) (string, int) {
	strs := strings.Split(step, " ")
	amount, _ := strconv.Atoi(strs[1])
	return strs[0], amount
}

func solvePart1(arr []string) {
	p := &positions{0, 0, 0}
	for _, step := range arr {
		direction, amount := splitStep(step)
		switch direction {
		case "forward":
			p.horizontal += amount
		case "down":
			p.depth += amount
		case "up":
			p.depth -= amount
		}
	}

	fmt.Println(p.horizontal * p.depth)
}

func solvePart2(arr []string) {
	p := &positions{0, 0, 0}
	for _, step := range arr {
		direction, amount := splitStep(step)
		switch direction {
		case "forward":
			p.horizontal += amount
			if p.aim > 0 {
				p.depth += (amount * p.aim)
			}
		case "down":
			p.aim += amount
		case "up":
			p.aim -= amount
		}
	}

	fmt.Println(p.horizontal * p.depth)
}
