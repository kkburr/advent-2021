package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	arr := setup()
	solvePart1(arr)
	solvePart2(arr)
}

func setup() []int {
	file, _ := os.Open("input")
	scanner := bufio.NewScanner(file)
	arr := make([]int, 0)
	for scanner.Scan() {
		v, _ := strconv.Atoi(scanner.Text())
		arr = append(arr, v)
	}
	return arr
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
