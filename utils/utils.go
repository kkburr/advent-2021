package utils

import (
	"bufio"
	"os"
	"strconv"
)

func Setup() []int {
	file, _ := os.Open("input")
	scanner := bufio.NewScanner(file)
	arr := make([]int, 0)
	for scanner.Scan() {
		v, _ := strconv.Atoi(scanner.Text())
		arr = append(arr, v)
	}
	return arr
}
