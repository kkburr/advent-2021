package main

//TODO: can I use pointers

import (
	"advent-2021/utils"
	"fmt"
	"strconv"
)

type game struct {
	boards       [][][]string
	winner       [][]string
	loser        [][]string
	drawnNumbers []string
	winnerSum    int
	loserSum     int
}

func main() {
	numbs, boards := utils.SetupDayFour()
	game := &game{
		boards,
		make([][]string, 0),
		make([][]string, 0),
		numbs,
		0,
		0,
	}
	game.solve()
}

func (g *game) solve() {
	for i := 0; i < len(g.drawnNumbers); i++ {
		numb := g.drawnNumbers[i]
		g.findMatches(numb)
		if i > 4 {
			g.findWinner()
			if len(g.winner) > 0 && g.winnerSum == 0 {
				n, _ := strconv.Atoi(numb)
				g.winnerSum = findSum(g.winner) * n
				fmt.Printf("Part 1 Answer: %v\n", g.winnerSum)
			} else if len(g.loser) > 0 {
				n, _ := strconv.Atoi(numb)
				g.loserSum = findSum(g.loser) * n
				fmt.Printf("Part 2 Answer: %v\n", g.loserSum)
				return
			}
		}
	}
}

func (g *game) findMatches(numb string) {
	for h := 0; h < len(g.boards); h++ {
		for i := 0; i < len(g.boards[h]); i++ {
			for j := 0; j < len(g.boards[h][i]); j++ {
				if g.boards[h][i][j] == numb {
					g.boards[h][i][j] = "X"
				}
			}
		}
	}
}

func (g *game) findWinner() {
	for h := 0; h < len(g.boards); h++ {
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				if len(g.boards) > h {
					if g.boards[h][i][j] == "X" {
						g.scan(h, i, j)
					}
				}

			}
		}
	}
}

func findSum(board [][]string) int {
	count := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			val := board[i][j]
			if val != "X" {
				n, _ := strconv.Atoi(val)
				count += n
			}
		}
	}
	return count
}

func (g *game) scan(i, rowIndex, colIndex int) {
	board := g.boards[i]
	row := board[rowIndex]
	colCount := 0
	rowCount := 0
	for j := 0; j < 5; j++ {
		if board[j][colIndex] == "X" {
			colCount += 1
		}
		if row[j] == "X" {
			rowCount += 1
		}
	}
	if rowCount == 5 || colCount == 5 {
		if len(g.winner) == 0 {
			g.winner = board
		} else if len(g.boards) == 1 {
			g.loser = board
		}
		l := len(g.boards) - 1
		g.boards[i] = g.boards[l]
		g.boards = g.boards[:l]
	}
}
