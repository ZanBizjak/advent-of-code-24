package day2

import (
	"fmt"

	"github.com/zanBizjak/advent-of-code-24/achelpers"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Task1() {
	reports := achelpers.IntReadGrid("day2.txt", " ")
	unsafe := 0
	prevLevel := 0

	for _, report := range reports {
		trend := 0
		prevTrend := 0
		for i, level := range report {
			if i == 0 {
				prevLevel = level
				continue
			}

			if level > prevLevel {
				trend = 1
			} else if level < prevLevel {
				trend = -1
			} else {
				trend = 0
			}
			if i == 1 {
				prevTrend = trend
				if trend == 0 {
					unsafe += 1
					break
				}
			}

			if trend != prevTrend {
				unsafe += 1
				break
			}

			if abs(level-prevLevel) > 3 {
				unsafe += 1
				break
			}

			prevTrend = trend
			prevLevel = level
		}
	}

	fmt.Println(len(reports) - unsafe)

}

func Task2() {

}
