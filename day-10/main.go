package main

import (
	"advent-2021/utils"
	"fmt"
	"sort"
)

var closers map[rune]rune = map[rune]rune{
	']': '[',
	'}': '{',
	'>': '<',
	')': '(',
}

var part1Score map[rune]int = map[rune]int{
	']': 57,
	'}': 1197,
	'>': 25137,
	')': 3,
}

var part2Score map[rune]int = map[rune]int{
	'[': 2,
	'{': 3,
	'<': 4,
	'(': 1,
}

func main() {
	arr := utils.SetupStringArray()
	part1 := 0
	part2 := []int{}
	for _, row := range arr {
		i := 0
		for {
			if i == len(row) {
				points := 0
				for i := len(row) - 1; i >= 0; i-- {
					points = (points * 5) + part2Score[rune(row[i])]
				}
				part2 = append(part2, points)
				break
			}
			char := rune(row[i])
			if val, ok := closers[char]; ok {
				if rune(row[i-1]) == val {
					row = row[:(i-1)] + row[(i+1):]
					i = i - 1
					continue
				} else {
					part1 += part1Score[char]
					break
				}
			}
			i += 1
		}
	}
	fmt.Printf("points: %v\n", part1)
	sort.Ints(part2)
	fmt.Printf("Part 2: %v\n", part2[len(part2)/2])
}
