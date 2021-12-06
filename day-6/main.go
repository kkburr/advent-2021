package main

import (
	"advent-2021/utils"
	"fmt"
)

func main() {
	arr := utils.SetupIntArray2()
	fmt.Println(solvePart1(arr))
	fmt.Println(solvePart2(arr))
}

func solvePart1(arr []int) int {
	for i := 0; i < 80; i++ {
		nextArr := make([]int, 0)
		arr = process(arr, nextArr)
	}
	return len(arr)
}

func solvePart2(arr []int) int {
	ages := map[int]int{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}
	for _, val := range arr {
		ages[val] = ages[val] + 1
	}
	for i := 0; i < 256; i++ {
		lastM := copyMap(ages)
		ages[8] = lastM[0]
		ages[7] = lastM[8]
		ages[6] = lastM[7] + lastM[0]
		ages[5] = lastM[6]
		ages[4] = lastM[5]
		ages[3] = lastM[4]
		ages[2] = lastM[3]
		ages[1] = lastM[2]
		ages[0] = lastM[1]
	}
	count := 0
	for _, v := range ages {
		count += v
	}
	return count
}

func process(firstArr, nextArr []int) []int {
	for _, v := range firstArr {
		if v > 0 {
			nextArr = append(nextArr, v-1)
		} else {
			nextArr = append(nextArr, 6, 8)
		}
	}
	return nextArr
}

func copyMap(m map[int]int) map[int]int {
	targetMap := make(map[int]int, 0)
	for key, value := range m {
		targetMap[key] = value
	}
	return targetMap
}
