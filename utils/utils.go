package utils

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getScanner() *bufio.Scanner {
	file, _ := os.Open("input")
	return bufio.NewScanner(file)
}

func SetupIntArray() []int {
	scanner := getScanner()
	arr := make([]int, 0)
	for scanner.Scan() {
		v, _ := strconv.Atoi(scanner.Text())
		arr = append(arr, v)
	}
	return arr
}

func SetupIntArray2() []int {
	//there's gotta be a better way to do this
	buff, _ := os.ReadFile("input")
	arr := strings.Split(strings.TrimSpace(string(buff)), ",")
	nextArr := make([]int, 0)
	for _, v := range arr {
		val, _ := strconv.Atoi(v)
		nextArr = append(nextArr, val)
	}
	return nextArr
}

func SetupStringArray() []string {
	// todo: simplify into something like:
	// buff, _ := os.ReadFile("input")
	// strings.Split(string(buff), "\n")
	scanner := getScanner()
	arr := make([]string, 0)
	for scanner.Scan() {
		row := scanner.Text()
		arr = append(arr, row)
	}
	return arr
}

func SetupDayFour() ([]string, [][][]string) {
	re := regexp.MustCompile(`\s+`)
	scanner := getScanner()
	numbs := make([]string, 0)
	boards := make([][][]string, 0)
	var nextBoard [][]string
	start := true
	for scanner.Scan() {
		row := scanner.Text()
		if start {
			numbs = strings.Split(row, ",")
			start = false
		} else if row == "" {
			if len(nextBoard) > 0 {
				boards = append(boards, nextBoard)
			}
			nextBoard = make([][]string, 0)
		} else {
			row = strings.TrimSpace(row)
			re.Split(row, -1)
			nextBoard = append(nextBoard, re.Split(row, -1))
		}
	}
	boards = append(boards, nextBoard)
	return numbs, boards
}
