package main

import (
	utils "advent-2021/utils"
	"fmt"
)

func main() {
	arr := utils.SetupIntArray()
	solvePart1(arr)
	solvePart2(arr)
}

func solvePart1(arr []int) {
	count := 0
	for i := 0; i < len(arr); i++ {
		if i > 0 && arr[i] > arr[i-1] {
			count++
		}
	}
	fmt.Println(count)
}

func solvePart2(arr []int) {
	count := 0
	lastSum := 0
	for i := 2; i < len(arr); i++ {
		nextSum := arr[i] + arr[i-1] + arr[i-2]
		if lastSum > 0 && nextSum > lastSum {
			count++
		}
		lastSum = nextSum
	}
	fmt.Println(count)
}
