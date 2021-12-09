package main

import (
	"advent-2021/utils"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var uniqueValsByLength = map[int]string{2: "1", 4: "4", 3: "7", 7: "8"}
var uniqueDigits = []int{2, 3, 4, 7}

func main() {
	solvePart1()
	solvePart2()
}

func solvePart1() {
	part1 := 0
	arr := utils.SetupStringArray2("|", " ")
	for _, val := range arr {
		if utils.Contains(uniqueDigits, len(val)) {
			part1 += 1
		}
	}
	fmt.Printf("Part 1: %v\n", part1)
}

func solvePart2() {
	part2 := 0
	arr := utils.SetupStringArray2b("|")
	for _, row := range arr {
		patterns := strings.Split(strings.TrimSpace(row[0]), " ")
		code := createCode(patterns)
		output := strings.Split(strings.TrimSpace(row[1]), " ")
		part2 += calculateOutput(output, code)
	}
	fmt.Printf("Part 2: %v\n", part2)
}

// 1 —> len == 2
// 4 —> len == 4
// 7 —> len == 3
// 8 —> len == 7
// 9 —> len == 6 && contains 4
// 0 —> len == 6 && contains 1
// 6 —> len == 6 && contains no unique numbers
// 3 —> len == 5 && contains 1
// 2 —> len == 5 && contains 2 letters in common with 4
// 5 —> len == 5 && contains 3 letters in common with 4

func createCode(arr []string) map[string]string {
	code := make(map[string]string, 0)
	var re4 *regexp.Regexp
	var re1 *regexp.Regexp
	for _, val := range arr {
		sorted, rex := processString(val)
		if len(sorted) == 4 {
			re4 = regexp.MustCompile(rex)
		} else if len(sorted) == 2 {
			re1 = regexp.MustCompile(rex)
		}
		length := len(val)
		if utils.Contains(uniqueDigits, length) {
			code[sorted] = uniqueValsByLength[length]
		} else {
			code[sorted] = ""
		}
	}
	for key, val := range code {
		if val == "" {
			matches4 := re4.FindAllString(key, -1)
			matches1 := re1.FindAllString(key, -1)
			if len(key) == 6 {
				if len(matches4) == 4 {
					code[key] = "9"
				} else if len(matches1) == 2 {
					code[key] = "0"
				} else {
					code[key] = "6"
				}
			} else if len(key) == 5 {
				if len(matches1) == 2 {
					code[key] = "3"
				} else if len(matches4) == 2 {
					code[key] = "2"
				} else if len(matches4) == 3 {
					code[key] = "5"
				}
			}
		}
	}
	return code
}

func processString(val string) (string, string) {
	s := splitAndSort(val)
	return strings.Join(s, ""), strings.Join(s, "|") // just do this for all of them (shrug)
}

func splitAndSort(val string) []string {
	s := strings.Split(val, "")
	sort.Strings(s)
	return s
}

func calculateOutput(arr []string, code map[string]string) int {
	temp := ""
	for _, val := range arr {
		s := splitAndSort(val)
		val = strings.Join(s, "")
		temp += code[val]
	}
	count, _ := strconv.Atoi(temp)
	return count
}
