package day5

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/zanBizjak/advent-of-code-24/achelpers"
)

func Task1() {
	before, after := achelpers.IntReadTwoColumns("day5rules.txt", "|")
	rules := make(map[int]map[int]bool, 0)
	for i, rule := range before {
		if rules[rule] == nil {
			rules[rule] = make(map[int]bool)
		}
		rules[rule][after[i]] = true
	}

	updates := achelpers.ReadRows("day5.txt")
	res := 0
	for _, update := range updates {
		res += inOrder(rules, update)

	}

	fmt.Println(res)
}

func inOrder(rules map[int]map[int]bool, update string) int {
	updateArr := strings.Split(update, ",")
	slices.Reverse(updateArr)
	var updateInt []int
	for _, num := range updateArr {
		i, _ := strconv.Atoi(num)
		updateInt = append(updateInt, i)
	}
	for i := 0; i < len(updateInt)-1; i++ {
		page := updateInt[i]
		for j := i + 1; j < len(updateInt); j++ {
			if rules[page][updateInt[j]] {
				return 0
			}

		}
	}

	return updateInt[len(updateInt)/2]
}

func Task2() {
}
