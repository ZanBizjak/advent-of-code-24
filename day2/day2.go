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

	for _, report := range reports {
		unsafety, _ := reportSafety(report)
		unsafe += unsafety

	}

	fmt.Println(len(reports) - unsafe)

}

func reportSafety(report []int) (int, int) {
	trend := 0
	prevLevel := 0
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
				return 1, i
			}
		}

		if trend != prevTrend {
			return 1, i
		}

		if abs(level-prevLevel) > 3 {
			return 1, i
		}

		prevTrend = trend
		prevLevel = level
	}
	return 0, -1
}

func Task2() {
	reports := achelpers.IntReadGrid("day2.txt", " ")
	unsafe := 0

	for _, report := range reports {
		unsafety, i := reportSafety(report)
		stillUnsafe := 1
		if unsafety == 1 {
			var oneRemoved []int
			oneRemoved = achelpers.IntRemoveIndex(report, 0)
			stillUnsafe, _ = reportSafety(oneRemoved)
			if stillUnsafe == 0 {
				continue
			}
			if i > 0 {
				oneRemoved = achelpers.IntRemoveIndex(report, i-1)
				stillUnsafe, _ = reportSafety(oneRemoved)
			}
			if stillUnsafe == 0 {
				continue
			}
			oneRemoved = achelpers.IntRemoveIndex(report, i)
			stillUnsafe, _ = reportSafety(oneRemoved)
			if stillUnsafe == 0 {
				continue
			}
			if i < len(report)-1 {
				oneRemoved = achelpers.IntRemoveIndex(report, i+1)
				stillUnsafe, _ = reportSafety(oneRemoved)
				if stillUnsafe == 0 {
					continue
				}

			}
		}
		unsafe += unsafety

	}

	fmt.Println(len(reports) - unsafe)
}
