package day1

import (
	"fmt"
	"sort"

	"github.com/zanBizjak/advent-of-code-24/achelpers"
)

func Task1() {
	col1, col2 := achelpers.IntReadTwoColumns("day1.txt", "   ")
	sort.Slice(col1, func(i, j int) bool { return col1[i] > col1[j] })
	sort.Slice(col2, func(i, j int) bool { return col2[i] > col2[j] })

	sum := 0
	for i := 0; i < len(col1); i++ {
		presum := col1[i] - col2[i]
		if presum < 0 {
			presum *= -1
		}
		sum += presum

	}

	fmt.Println(sum)
}

func Task2() {
	left, right := achelpers.IntReadTwoColumns("day1.txt", "   ")
	sort.Slice(left, func(i, j int) bool { return left[i] > left[j] })
	sort.Slice(right, func(i, j int) bool { return right[i] > right[j] })

	rightReps := make(map[int]int)

	for _, val := range right {
		rightReps[val] = rightReps[val] + 1
	}
	sum := 0
	for _, val := range left {
		sum += val * rightReps[val]
	}

	fmt.Println(sum)

}
