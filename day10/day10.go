package day10

import (
	"fmt"

	"github.com/zanBizjak/advent-of-code-24/achelpers"
)

type Pos struct {
	x int
	y int
}

func Task1() {
	grid := achelpers.IntReadGrid("day10.txt", "")
	res := 0
	for i, row := range grid {
		for j, pos := range row {
			if pos == 0 {
				res += trailheadScoreP1(grid, i, j, pos, 0, make(map[Pos]bool))
			}
		}
	}

	fmt.Println(res)
}

func trailheadScoreP1(grid [][]int, y, x, pos, score int, found map[Pos]bool) int {
	if pos == 9 {
		if found[Pos{x: x, y: y}] {
			return 0
		}
		found[Pos{x: x, y: y}] = true
		return 1
	}
	if y-1 >= 0 {
		if grid[y-1][x] == pos+1 {

			score += trailheadScoreP1(grid, y-1, x, grid[y-1][x], 0, found)
		}
	}
	if y+1 < len(grid) {
		if grid[y+1][x] == pos+1 {
			score += trailheadScoreP1(grid, y+1, x, grid[y+1][x], 0, found)
		}
	}
	if x-1 >= 0 {
		if grid[y][x-1] == pos+1 {
			score += trailheadScoreP1(grid, y, x-1, grid[y][x-1], 0, found)
		}
	}
	if x+1 < len(grid) {
		if grid[y][x+1] == pos+1 {
			score += trailheadScoreP1(grid, y, x+1, grid[y][x+1], 0, found)
		}
	}

	return score
}

func trailheadScoreP2(grid [][]int, y, x, pos, score int) int {
	if pos == 9 {
		return 1
	}
	if y-1 >= 0 {
		if grid[y-1][x] == pos+1 {

			score += trailheadScoreP2(grid, y-1, x, grid[y-1][x], 0)
		}
	}
	if y+1 < len(grid) {
		if grid[y+1][x] == pos+1 {
			score += trailheadScoreP2(grid, y+1, x, grid[y+1][x], 0)
		}
	}
	if x-1 >= 0 {
		if grid[y][x-1] == pos+1 {
			score += trailheadScoreP2(grid, y, x-1, grid[y][x-1], 0)
		}
	}
	if x+1 < len(grid) {
		if grid[y][x+1] == pos+1 {
			score += trailheadScoreP2(grid, y, x+1, grid[y][x+1], 0)
		}
	}

	return score
}

func Task2() {

	grid := achelpers.IntReadGrid("day10.txt", "")
	res := 0
	for i, row := range grid {
		for j, pos := range row {
			if pos == 0 {
				res += trailheadScoreP2(grid, i, j, pos, 0)
			}
		}
	}

	fmt.Println(res)
}
