package day4

import (
	"fmt"
	"slices"

	"github.com/zanBizjak/advent-of-code-24/achelpers"
)

func Task1() {
	grid := achelpers.RuneGetGrid("day4.txt")
	numOfXmas := 0
	for _, words := range grid {
		numOfXmas += findXmas(words)
	}

	for _, words := range grid {
		reversed := achelpers.RuneCopySlice(words)
		slices.Reverse(reversed)
		numOfXmas += findXmas(words)
	}

	for i := len(grid) - 1; i >= 0; i-- {
		words := getDiagonal(grid, i, 0, 1, 1)
		numOfXmas += findXmas(words)
		words = getDiagonal(grid, 0, i, 1, 1)
		numOfXmas += findXmas(words)
		words = getDiagonal(grid, len(grid)-1, i, -1, -1)
		numOfXmas += findXmas(words)
		words = getDiagonal(grid, i, len(grid)-1, -1, -1)
		numOfXmas += findXmas(words)

		words = getDiagonal(grid, len(grid)-1, i, -1, 1)
		numOfXmas += findXmas(words)
		words = getDiagonal(grid, i, len(grid)-1, -1, 1)
		numOfXmas += findXmas(words)

		words = getDiagonal(grid, 0, i, 1, 0)
		numOfXmas += findXmas(words)
		words = getDiagonal(grid, len(grid)-1, i, -1, 0)
		numOfXmas += findXmas(words)

		words = getDiagonal(grid, i, 0, 0, 1)
		numOfXmas += findXmas(words)
		words = getDiagonal(grid, i, len(grid)-1, 0, -1)
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
				runeState++
				break
			}
			runeState = 0
		case 1:
			if char == 'M' {
				runeState++
				break
			}
			runeState = 0

		case 2:
			if char == 'A' {
				runeState++
				break
			}
			runeState = 0
		case 3:
			if char == 'S' {
				runeState++
				break
			}
			runeState = 0
		}
		if runeState == 4 {
			numOfXmas++
			runeState = 0
		}

	}

	return numOfXmas
}

func Task2() {
}
