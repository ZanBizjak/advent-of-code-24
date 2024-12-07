package day7

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/zanBizjak/advent-of-code-24/achelpers"
)

func Task1() {
	inputs := achelpers.ReadRows("day7.txt")
	var res int64 = 0
	for _, input := range inputs {
		res += validResultv2(input, 2)
	}
	fmt.Println(res)
}

func validResultv2(input string, numOfOps int) int64 {
	resAndNums := strings.Split(input, ": ")
	result, err := strconv.ParseInt(resAndNums[0], 10, 64)

	if err != nil {
		log.Fatal("error converting the result portion of the input")
	}
	nums := achelpers.StrToIntSlice(resAndNums[1], " ")
	combos := pow(numOfOps, len(nums)-1)
	opTable := make([]int, len(nums)-1)
	for opFlags := 0; opFlags < combos; opFlags++ {

		copyNums := achelpers.IntCopySlice(nums)
		for i := 0; i < len(opTable); i++ {
			op := opTable[i]
			var together int64
			if op == 0 {
				together = copyNums[0] + copyNums[1]
			} else if op == 1 {
				together = copyNums[0] * copyNums[1]
			} else if op == 2 {
				cc, err := strconv.ParseInt(fmt.Sprintf("%v%v", copyNums[0], copyNums[1]), 10, 64)
				if err != nil {
					log.Panic("AAAAAAAA")
				}
				together = cc

			}

			copyNums = copyNums[1:]
			copyNums[0] = together
		}
		if copyNums[0] == result {
			return result
		}
		opTable = increment(opTable, 0, numOfOps)
	}

	return 0
}

func pow(base int, power int) int {
	ret := 1
	for i := 0; i < power; i++ {
		ret *= base
	}

	return ret

}

func increment(incTable []int, toInc int, max int) []int {
	if toInc >= len(incTable) {
		return incTable
	}
	incTable[toInc] += 1
	if incTable[toInc] >= max {
		incTable[toInc] = 0
		return increment(incTable, toInc+1, max)
	}

	return incTable
}

func Task2() {
	inputs := achelpers.ReadRows("day7.txt")
	var res int64 = 0
	for _, input := range inputs {
		res += validResultv2(input, 3)
	}
	fmt.Println(res)
}
