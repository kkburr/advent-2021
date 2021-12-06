package main

import (
	"advent-2021/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type coordinates struct {
	x int
	y int
}

type graph struct {
	points map[coordinates]int
}

func main() {
	vents := utils.SetupStringArray()
	g := &graph{
		make(map[coordinates]int, 0),
	}
	g.processCoordinates(vents)
}

func (g *graph) processCoordinates(arr []string) {
	for _, val := range arr {
		strs := strings.Split(val, " -> ")
		start := strings.Split(strs[0], ",")
		end := strings.Split(strs[1], ",")
		if start[0] == end[0] || start[1] == end[1] {
			g.storeHVCoordinates(start, end)
		} else {
			x1, _ := strconv.Atoi(start[0])
			y1, _ := strconv.Atoi(start[1])
			x2, _ := strconv.Atoi(end[0])
			y2, _ := strconv.Atoi(end[1])
			if (x1 == y1 && x2 == y2) || math.Abs(float64(x1-x2)) == math.Abs(float64(y1-y2)) {
				g.storeDCoordinates(x1, y1, x2, y2)
			}
		}
	}
	g.findTwoMatches()
}

func (g *graph) storeDCoordinates(x1, y1, x2, y2 int) {
	var currentX int
	var currentY int
	var stopX int
	var stopY int

	if x1 < x2 {
		currentX = x1
		currentY = y1
		stopX = x2
		stopY = y2
	} else {
		currentX = x2
		currentY = y2
		stopX = x1
		stopY = y1
	}

	for {
		coords := coordinates{currentX, currentY}
		if g.points[coords] > 0 {
			g.points[coords] = g.points[coords] + 1
		} else {
			g.points[coords] = 1
		}
		if currentX == stopX {
			return
		}
		currentX += 1
		if currentY < stopY {
			currentY += 1
		} else {
			currentY -= 1
		}
	}
}
func (g *graph) storeHVCoordinates(start, end []string) {
	x1, _ := strconv.Atoi(start[0])
	y1, _ := strconv.Atoi(start[1])
	x2, _ := strconv.Atoi(end[0])
	y2, _ := strconv.Atoi(end[1])
	var current int
	var stop int
	if x1 != x2 {
		if x1 < x2 {
			current = x1
			stop = x2
		} else {
			current = x2
			stop = x1
		}
		for {
			coords := coordinates{current, y2}
			if g.points[coords] > 0 {
				g.points[coords] = g.points[coords] + 1
			} else {
				g.points[coords] = 1
			}
			if current == stop {
				return
			}
			current += 1
		}
	}

	if y1 != y2 {
		if y1 < y2 {
			current = y1
			stop = y2
		} else {
			current = y2
			stop = y1
		}
		for {
			coords := coordinates{x1, current}
			if g.points[coords] > 0 {
				g.points[coords] = g.points[coords] + 1
			} else {
				g.points[coords] = 1
			}
			if current == stop {
				return
			}
			current += 1
		}
	}
}

func (g *graph) findTwoMatches() {
	counter := 0
	for _, v := range g.points {
		if v > 1 {
			counter++
		}
	}
	fmt.Printf("answer: %v\n", counter)
}
