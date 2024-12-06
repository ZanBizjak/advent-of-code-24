package day6

import (
	"fmt"

	"github.com/zanBizjak/advent-of-code-24/achelpers"
)

func Task1() {
	grid := achelpers.RuneGetGrid("day6.txt")

	y, x := findGuard(grid)
	uniqueSteps(grid, x, y)
	res := 0
	for _, row := range grid {
		for _, pos := range row {
			if pos == 'X' {
				res++
			}

		}
	}

	for _, row := range grid {
		for _, pos := range row {
			fmt.Print(string(pos))
		}
		fmt.Println()
	}
	fmt.Println(res)

}

func uniqueSteps(grid [][]rune, x int, y int) {
	for x < len(grid) && y < len(grid) && x >= 0 && y >= 0 {
		guardFace := grid[y][x]
		grid[y][x] = 'X'
		x, y, guardFace = step(grid, x, y, guardFace)
		if guardFace != 'Q' {
			grid[y][x] = guardFace
		}
	}
}

func step(grid [][]rune, x, y int, guardFace rune) (int, int, rune) {
	plusX := 0
	plusY := 0
	newFace := ' '
	switch guardFace {
	case '^':
		if y-1 < 0 {
			return x, y - 1, 'Q'
		}
		if grid[y-1][x] == '#' {
			if x+1 >= len(grid) {
				return x + 1, y, 'Q'
			}
			newFace = '>'
			plusX = 1
		} else {
			newFace = '^'
			plusY = -1
		}

	case '>':
		if x+1 >= len(grid) {
			return x + 1, y, 'Q'
		}
		if grid[y][x+1] == '#' {
			if y+1 >= len(grid) {
				return x, y + 1, 'Q'
			}
			newFace = 'ˇ'
			plusY = 1
		} else {
			newFace = '>'
			plusX = 1
		}

	case 'ˇ':
		if y+1 >= len(grid) {
			return x, y + 1, 'Q'
		}
		if grid[y+1][x] == '#' {
			if x-1 >= len(grid) {
				return x - 1, y, 'Q'
			}
			newFace = '<'
			plusX = -1
		} else {
			newFace = 'ˇ'
			plusY = 1
		}
	case '<':
		if x-1 < 0 {
			return x - 1, y, 'Q'
		}
		if grid[y][x-1] == '#' {
			if y-1 >= len(grid) {
				return x, y - 1, 'Q'
			}
			newFace = '^'
			plusY = -1
		} else {
			newFace = '<'
			plusX = -1
		}
	}
	if grid[y+plusY][x+plusX] == '#' {
		return step(grid, x, y, newFace)
	}

	return x + plusX, y + plusY, newFace

}

func findGuard(grid [][]rune) (int, int) {
	for y, row := range grid {
		for x, point := range row {
			if point == '^' || point == '>' || point == 'ˇ' || point == '<' {
				return y, x

			}
		}

	}

	return -1, -1

}

func Task2() {
	grid := achelpers.RuneGetGrid("day6.txt")

	y, x := findGuard(grid)
	uniqueStepGrid := achelpers.RuneCopyGrid(grid)
	uniqueSteps(uniqueStepGrid, x, y)

	res := 0
	for obstacleY, row := range uniqueStepGrid {
		for obstacleX, pos := range row {
			if pos == 'X' {
				if obstacleX == x && obstacleY == y {
					continue
				}
				newObstacleGrid := achelpers.RuneCopyGrid(grid)
				newObstacleGrid[obstacleY][obstacleX] = '#'
				if isCycling(newObstacleGrid, x, y) {
					res++
				}

			}

		}
	}

	fmt.Println(res)
}

func isCycling(grid [][]rune, x, y int) bool {
	for x < len(grid) && y < len(grid) && x >= 0 && y >= 0 {
		guardFace := grid[y][x]
		xStep, yStep, newFace := step(grid, x, y, guardFace)

		if newFace == 'Q' {
			// not cycling
			return false
		}

		if grid[yStep][xStep] == newFace {
			// cycling
			return true
		}
		x, y = xStep, yStep
		grid[y][x] = newFace

	}
	return false
}
