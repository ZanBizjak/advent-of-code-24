package day6

import (
	"fmt"

	"github.com/zanBizjak/advent-of-code-24/achelpers"
)

func Task1() {
	grid := achelpers.RuneGetGrid("day6.txt")

	y, x := findGuard(grid)
	for x < len(grid) && y < len(grid) {
		guardFace := grid[y][x]
		grid[y][x] = 'X'
		x, y = step(grid, x, y, guardFace)
	}
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

func step(grid [][]rune, x, y int, guardFace rune) (int, int) {
	plusX := 0
	plusY := 0
	switch guardFace {
	case '^':
		if y-1 < 0 {
			return x, y - 1
		}
		if grid[y-1][x] == '#' {
			if x+1 >= len(grid) {
				return x + 1, y
			}
			grid[y][x+1] = '>'
			plusX = 1
		} else {
			grid[y-1][x] = '^'
			plusY = -1
		}

	case '>':
		if x+1 >= len(grid) {
			return x + 1, y
		}
		if grid[y][x+1] == '#' {
			if y+1 >= len(grid) {
				return x, y + 1
			}
			grid[y+1][x] = 'ˇ'
			plusY = 1
		} else {
			grid[y][x+1] = '>'
			plusX = 1
		}

	case 'ˇ':
		if y+1 >= len(grid) {
			return x, y + 1
		}
		if grid[y+1][x] == '#' {
			if x-1 >= len(grid) {
				return x - 1, y
			}
			grid[y][x-1] = '<'
			plusX = -1
		} else {
			grid[y+1][x] = 'ˇ'
			plusY = 1
		}
	case '<':
		if x-1 >= len(grid) {
			return x - 1, y
		}
		if grid[y][x-1] == '#' {
			if y-1 >= len(grid) {
				return x, y - 1
			}
			grid[y-1][x] = '^'
			plusY = -1
		} else {
			grid[y][x-1] = '<'
			plusX = -1
		}
	}

	return x + plusX, y + plusY

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
}
