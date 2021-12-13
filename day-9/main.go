package main

import (
	"advent-2021/utils"
	"fmt"
	"sort"
)

type coord struct {
	y int
	x int
}

type caves struct {
	tubes     [][]int
	basins    []int
	lowPoints []coord
}

func main() {
	arr := utils.SetupIntArray3()
	c := &caves{
		arr,
		[]int{},
		[]coord{},
	}
	count := c.solvePart1()
	fmt.Printf("Part 1: %v\n", count)
	count = c.solvePart2()
	fmt.Printf("Part 2: %v\n", count)
}

func (c *caves) solvePart1() int {
	count := 0
	var lowPoints = []int{}
	for y, row := range c.tubes {
		for x, val := range row {
			if c.isLowPoint(x, y) {
				lowPoints = append(lowPoints, val)
				c.lowPoints = append(c.lowPoints, coord{y, x})
			}
		}
	}
	for _, val := range lowPoints {
		count += (val + 1)
	}
	return count
}

func (c *caves) isLowPoint(x, y int) bool {
	var adjacentVals = []int{}
	if y > 0 {
		adjacentVals = append(adjacentVals, c.tubes[y-1][x])
	}
	if y < len(c.tubes)-1 {
		adjacentVals = append(adjacentVals, c.tubes[y+1][x])
	}
	if x > 0 {
		adjacentVals = append(adjacentVals, c.tubes[y][x-1])
	}
	if x < len(c.tubes[y])-1 {
		adjacentVals = append(adjacentVals, c.tubes[y][x+1])
	}
	sort.Ints(adjacentVals)
	if c.tubes[y][x] < adjacentVals[0] {
		return true
	}
	return false
}

func (c *caves) solvePart2() int {
	for _, coordinate := range c.lowPoints {
		x := coordinate.x
		y := coordinate.y
		traversed := make(map[coord]struct{}, 0)
		basinSize, _ := c.recurse(x, y, 0, traversed)
		if basinSize > 0 {
			c.basins = append(c.basins, basinSize)
		}
	}
	sort.Ints(c.basins)
	l := len(c.basins)
	return c.basins[l-1] * c.basins[l-2] * c.basins[l-3] // currently getting 803225. correct answer is 1148965
}

func (c *caves) recurse(x, y, size int, traversed map[coord]struct{}) (int, map[coord]struct{}) {
	coo := coord{y, x}
	traversed[coo] = struct{}{}

	val := c.tubes[y][x]
	if val == 9 {
		return size, traversed
	}
	size += 1
	if y > 0 {
		next := coord{y - 1, x}
		_, ok := traversed[next]
		if val+1 == c.tubes[y-1][x] && !ok {
			size, traversed = c.recurse(x, y-1, size, traversed)
		}
	}
	if y < len(c.tubes)-1 {
		next := coord{y + 1, x}
		_, ok := traversed[next]
		if val+1 == c.tubes[y+1][x] && !ok {
			size, traversed = c.recurse(x, y+1, size, traversed)
		}
	}
	if x > 0 {
		next := coord{y, x - 1}
		_, ok := traversed[next]
		if val+1 == c.tubes[y][x-1] && !ok {
			size, traversed = c.recurse(x-1, y, size, traversed)
		}
	}
	if x < len(c.tubes[y])-1 {
		next := coord{y, x + 1}
		_, ok := traversed[next]
		if val+1 == c.tubes[y][x+1] && !ok {
			size, traversed = c.recurse(x+1, y, size, traversed)
		}
	}
	return size, traversed

}
