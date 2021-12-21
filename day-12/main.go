package main

import (
	"advent-2021/utils"
	"fmt"
	"strings"
	"unicode"
)

type cave struct {
	routes map[string][]string
}

func main() {
	arr := utils.SetupStringArray()
	routes := setupRoutes(arr)
	c := &cave{routes}
	part1 := c.part1("start", "start", []string{"start"}, []string{})
	fmt.Printf("Part one answer: %v\n", len(part1))
	part2 := c.part2("start", "start", []string{"start"}, []string{}, true)
	fmt.Printf("Part two answer: %v\n", len(part2))
}

func setupRoutes(rows []string) map[string][]string {
	routes := make(map[string][]string)
	for _, row := range rows {
		arr := strings.Split(row, "-")
		j := 1
		for i, loc := range arr {
			if i == 1 {
				j = 0
			}
			if connections, ok := routes[loc]; ok {
				routes[loc] = append(connections, arr[j])
			} else {
				routes[loc] = []string{arr[j]}
			}
		}
	}
	return routes
}

func (c *cave) part1(cave, route string, traveled, routes []string) []string {
	if cave == "end" {
		return append(routes, route)
	}
	if cave == "start" {
		for _, next := range c.routes[cave] {
			nextRoute := route + "," + next
			nextTraveled := append(traveled, next)
			nextRoutes := c.part1(next, nextRoute, nextTraveled, []string{})
			routes = append(routes, nextRoutes...)
		}
		return routes
	}
	for _, next := range c.routes[cave] {
		if unicode.IsLower(rune(next[0])) && utils.ContainsString(traveled, next) {
			continue
		} else {
			nextRoute := route + "," + next
			nextTraveled := append(traveled, next)
			routes = c.part1(next, nextRoute, nextTraveled, routes)
		}
	}
	return routes
}

func (c *cave) part2(cave, route string, traveled, routes []string, first bool) []string {
	if cave == "end" {
		return append(routes, route)
	}
	if cave == "start" {
		for _, next := range c.routes[cave] {
			nextRoute := route + "," + next
			nextTraveled := append(traveled, next)
			nextRoutes := c.part2(next, nextRoute, nextTraveled, []string{}, first)
			routes = append(routes, nextRoutes...)
		}
		return routes
	}
	for _, next := range c.routes[cave] {
		if unicode.IsLower(rune(next[0])) && utils.ContainsString(traveled, next) {
			if first && next != "start" {
				nextRoute := route + "," + next
				nextTraveled := append(traveled, next)
				routes = c.part2(next, nextRoute, nextTraveled, routes, false)
			}
			continue
		} else {
			nextRoute := route + "," + next
			nextTraveled := append(traveled, next)
			routes = c.part2(next, nextRoute, nextTraveled, routes, first)
		}
	}
	return routes
}
