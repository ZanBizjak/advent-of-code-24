package day3

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/zanBizjak/advent-of-code-24/achelpers"
)

func Task1() {
	input := achelpers.StringReadFile("day3.txt")
	fmt.Println(mul(input))

}

func mul(input string) int {
	matcher := 0
	firstLen := 0
	firstIndex := 0
	secLen := 0
	secondIndex := 0
	result := 0
	for i := 0; i < len(input); i++ {
		if match(string(input[i]), matcher) {
			if matcher == 4 || matcher == 6 {
				numLen := findNumLen(input, i, 0)
				if numLen > 3 {
					matcher = 0
					continue
				}
				if matcher == 4 {
					firstLen = numLen
					firstIndex = i
				} else {
					secLen = numLen
					secondIndex = i
				}
				i += numLen - 1
			}
			matcher++
		} else {
			matcher = 0
		}
		if matcher == 8 {

			first, _ := strconv.Atoi(string(input[firstIndex : firstIndex+firstLen]))
			second, _ := strconv.Atoi(string(input[secondIndex : secondIndex+secLen]))
			result += first * second
			matcher = 0

		}

	}

	return result
}

func findNumLen(input string, i, len int) int {
	if _, err := strconv.Atoi(string(input[i])); err == nil {
		return findNumLen(input, i+1, len+1)
	}

	return len
}

func match(s string, matcher int) bool {
	switch matcher {
	case 0:
		if s == "m" {
			return true
		}
	case 1:
		if s == "u" {
			return true
		}
	case 2:
		if s == "l" {
			return true
		}
	case 3:
		if s == "(" {
			return true
		}
	case 4:
		if _, err := strconv.Atoi(s); err == nil {
			return true

		}
	case 5:
		if s == "," {
			return true
		}
	case 6:
		if _, err := strconv.Atoi(s); err == nil {
			return true
		}
	case 7:
		if s == ")" {
			return true
		}
	}

	return false
}
func Task2() {
	input := achelpers.StringReadFile("day3.txt")

	result := 0
	dos := strings.Split(input, "do()")
	for _, doString := range dos {
		donts := strings.Split(doString, "don't()")
		result += mul(donts[0])

	}

	fmt.Println(result)

}
