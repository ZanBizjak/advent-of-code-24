package day8

import (
	"fmt"

	"github.com/zanBizjak/advent-of-code-24/achelpers"
)

type NodePos struct {
	x int
	y int
}

func Task1() {
	grid := achelpers.RuneGetGrid("day8.txt")
	var nodes map[rune][]NodePos = make(map[rune][]NodePos)

	for y, row := range grid {
		for x, node := range row {
			if node != '.' {
				nodes[node] = append(nodes[node], NodePos{x: x, y: y})
			}
		}
	}

	var uniqueFrequencies map[NodePos]bool = make(map[NodePos]bool)
	for _, antennaPos := range nodes {
		findAntinodes(antennaPos, len(grid), uniqueFrequencies)
	}

	fmt.Println(len(uniqueFrequencies))

}

func findAntinodes(antennaPos []NodePos, maxLen int, found map[NodePos]bool) {
	for i := 0; i < len(antennaPos)-1; i++ {
		firstAntenna := antennaPos[i]
		for j := i + 1; j < len(antennaPos); j++ {
			secondAntenna := antennaPos[j]
			x, y := findDistance(firstAntenna, secondAntenna)
			x = firstAntenna.x - (x * 2)
			y = firstAntenna.y - (y * 2)
			if x >= 0 && x < maxLen && y >= 0 && y < maxLen {
				found[NodePos{x: x, y: y}] = true
			}
			x, y = findDistance(secondAntenna, firstAntenna)
			x = secondAntenna.x - (x * 2)
			y = secondAntenna.y - (y * 2)
			if x >= 0 && x < maxLen && y >= 0 && y < maxLen {
				found[NodePos{x: x, y: y}] = true
			}

		}
	}

}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func findDistance(firstAntenna, secondAntenna NodePos) (int, int) {
	x := (firstAntenna.x - secondAntenna.x)
	y := (firstAntenna.y - secondAntenna.y)

	return x, y

}

func Task2() {
	grid := achelpers.RuneGetGrid("day8.txt")
	var nodes map[rune][]NodePos = make(map[rune][]NodePos)

	for y, row := range grid {
		for x, node := range row {
			if node != '.' {
				nodes[node] = append(nodes[node], NodePos{x: x, y: y})
			}
		}
	}

	var uniqueFrequencies map[NodePos]bool = make(map[NodePos]bool)
	for _, antennaPos := range nodes {
		findContinuesAntinodes(antennaPos, len(grid), uniqueFrequencies)
	}

	fmt.Println(len(uniqueFrequencies))
}

func findContinuesAntinodes(antennaPos []NodePos, maxLen int, found map[NodePos]bool) {
	for i := 0; i < len(antennaPos)-1; i++ {
		firstAntenna := antennaPos[i]
		for j := i + 1; j < len(antennaPos); j++ {
			secondAntenna := antennaPos[j]
			x, y := findDistance(firstAntenna, secondAntenna)
			continuousAntiNode(firstAntenna, maxLen, found, x, y)
			x, y = findDistance(secondAntenna, firstAntenna)
			continuousAntiNode(secondAntenna, maxLen, found, x, y)
		}
	}
}

func continuousAntiNode(node NodePos, maxLen int, found map[NodePos]bool, x, y int) {
	node.x = node.x - x
	node.y = node.y - y
	if node.x >= 0 && node.x < maxLen && node.y >= 0 && node.y < maxLen {
		found[node] = true

		continuousAntiNode(node, maxLen, found, x, y)
	}
	return

}
