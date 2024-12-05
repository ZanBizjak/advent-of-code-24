package day4

import (
	"fmt"
	"slices"

	"github.com/zanBizjak/advent-of-code-24/achelpers"
)

func Task1() {
	grid := achelpers.RuneGetGrid("day4.txt")
	numOfXmas := 0
	for i := 0; i < len(grid); i++ {
		words := getDiagonal(grid, i, 0, 1, 1)
		numOfXmas += findXmas(words)
		slices.Reverse(words)
		numOfXmas += findXmas(words)
		if i != 0 {
			words = getDiagonal(grid, 0, i, 1, 1)
			numOfXmas += findXmas(words)
			slices.Reverse(words)
			numOfXmas += findXmas(words)
		}

		words = getDiagonal(grid, len(grid)-1, i, -1, 1)
		numOfXmas += findXmas(words)
		slices.Reverse(words)
		numOfXmas += findXmas(words)
		if i != len(grid)-1 {
			words = getDiagonal(grid, i, 0, -1, 1)
			numOfXmas += findXmas(words)
			slices.Reverse(words)
			numOfXmas += findXmas(words)
		}

		words = getDiagonal(grid, i, 0, 0, 1)
		numOfXmas += findXmas(words)
		slices.Reverse(words)
		numOfXmas += findXmas(words)

		words = getDiagonal(grid, 0, i, 1, 0)
		numOfXmas += findXmas(words)
		slices.Reverse(words)
		numOfXmas += findXmas(words)

	}

	fmt.Println(numOfXmas)

}

func getDiagonal(grid [][]rune, startY, startX, dirY, dirX int) []rune {
	var word []rune
	xPos := startX
	yPos := startY

	for yPos < len(grid) && yPos >= 0 && xPos < len(grid) && xPos >= 0 {
		word = append(word, grid[yPos][xPos])
		fmt.Print(string(grid[yPos][xPos]))
		xPos += dirX
		yPos += dirY

	}

	return word
}

func findXmas(r []rune) int {
	runeState := 0
	numOfXmas := 0
	for _, char := range r {
		switch runeState {
		case 0:
			if char == 'X' {
				runeState += 1
				break
			}
			runeState = 0
		case 1:
			if char == 'M' {
				runeState += 1
				break
			}
			if char == 'X' {
				runeState = 1
			} else {
				runeState = 0
			}

		case 2:
			if char == 'A' {
				runeState += 1
				break
			}
			if char == 'X' {
				runeState = 1
			} else {
				runeState = 0
			}
		case 3:
			if char == 'S' {
				runeState += 1
				break
			}
			if char == 'X' {
				runeState = 1
			} else {
				runeState = 0
			}
		}
		if runeState == 4 {
			numOfXmas++
			runeState = 0
		}

	}

	return numOfXmas
}

func Task2() {
	grid := achelpers.RuneGetGrid("day4.txt")
	numOfMas := 0
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid)-1; j++ {
			if grid[i][j] == 'A' {
				if (grid[i-1][j-1] == 'M' && grid[i+1][j+1] == 'S') || (grid[i-1][j-1] == 'S' && grid[i+1][j+1] == 'M') {

					if (grid[i+1][j-1] == 'M' && grid[i-1][j+1] == 'S') || (grid[i+1][j-1] == 'S' && grid[i-1][j+1] == 'M') {
						numOfMas++
					}
				}
			}

		}
	}
	fmt.Println(numOfMas)
}
