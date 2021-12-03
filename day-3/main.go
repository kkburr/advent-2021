package main

import (
	"advent-2021/utils"
	"fmt"
	"strconv"
)

func main() {
	arr := utils.SetupStringArray()
	solvePart1(arr)
	solvePart2(arr)
}

func solvePart1(arr []string) {
	var gamma string
	var epsilon string
	for i := 0; i < len(arr[0]); i++ {
		count0 := 0
		count1 := 0
		for j := 0; j < len(arr); j++ {
			if arr[j][i] == '0' {
				count0++
			} else {
				count1++
			}
		}
		if count0 > count1 {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}
	gammaRate, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonRate, _ := strconv.ParseInt(epsilon, 2, 64)
	fmt.Printf("Part 1 answer: %d\n", gammaRate*epsilonRate)
}

func solvePart2(arr []string) {
	oxyRatingStr := getRating(arr, 0, true)[0]
	co2RatingStr := getRating(arr, 0, false)[0]
	oxyRating, _ := strconv.ParseInt(oxyRatingStr, 2, 64)
	co2Rating, _ := strconv.ParseInt(co2RatingStr, 2, 64)
	fmt.Printf("Part 2 answer: %d\n", oxyRating*co2Rating)
}

func getRating(arr []string, pos int, oxy bool) []string {
	nextArr := []string{}
	var keep rune
	count0 := 0
	count1 := 0
	for i := 0; i < len(arr); i++ {
		if arr[i][pos] == '0' {
			count0++
		} else {
			count1++
		}
	}
	if (oxy && count1 >= count0) || (!oxy && count0 > count1) {
		keep = '1'
	} else {
		keep = '0'
	}
	for i := 0; i < len(arr); i++ {
		if arr[i][pos] == byte(keep) {
			nextArr = append(nextArr, arr[i])
		}
	}
	if len(nextArr) > 1 {
		return getRating(nextArr, pos+1, oxy)
	}
	return nextArr
}
