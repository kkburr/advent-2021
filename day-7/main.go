package main

import (
	"advent-2021/utils"
	"fmt"
	"math"
	"sort"

	"github.com/montanaflynn/stats"
)

func main() {
	arr := utils.SetupFloatArray()
	median, _ := stats.Median(arr)
	part1 := 0
	for _, val := range arr {
		part1 += int(math.Abs(val - median))
	}
	fmt.Printf("Part 1: %v\n", part1)

	// part 2
	mean, _ := stats.Mean(arr)
	rounded := math.Round(mean)
	//probably not the best logic but it got me the right answer
	if rounded-mean < .25 {
		mean = rounded
	} else {
		mean = rounded - 1
	}

	m := createMap(arr, mean)
	part2 := 0

	for _, val := range arr {
		part2 += m[int(math.Abs(val-mean))]
	}
	fmt.Printf("Part 2: %v\n", part2)
}

func createMap(arr []float64, mean float64) map[int]int {
	var diff float64
	sort.Float64s(arr)
	max := arr[len(arr)-1]
	min := arr[0]
	if max-mean > mean-min {
		diff = max - mean
	} else {
		diff = mean - min
	}
	m := make(map[int]int, 0)
	counter := 0
	for i := 1; i <= int(diff); i++ {
		counter = i + counter
		m[i] = counter
	}
	return m
}
