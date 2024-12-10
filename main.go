package main

import (
	"os"

	"github.com/zanBizjak/advent-of-code-24/day1"
	"github.com/zanBizjak/advent-of-code-24/day2"
	"github.com/zanBizjak/advent-of-code-24/day3"
	"github.com/zanBizjak/advent-of-code-24/day4"
	"github.com/zanBizjak/advent-of-code-24/day5"
	"github.com/zanBizjak/advent-of-code-24/day6"
	"github.com/zanBizjak/advent-of-code-24/day7"
	"github.com/zanBizjak/advent-of-code-24/day8"
	"github.com/zanBizjak/advent-of-code-24/day9"
)

func main() {
	day := os.Args[1]
	task := os.Args[2]

	switch day {
	case "1":
		taskDay1(task)
	case "2":
		taskDay2(task)
	case "3":
		taskDay3(task)
	case "4":
		taskDay4(task)
	case "5":
		taskDay5(task)
	case "6":
		taskDay6(task)
	case "7":
		taskDay7(task)
	case "8":
		taskDay8(task)
	case "9":
		taskDay9(task)
	}

}

func taskDay1(task string) {
	if task == "1" {
		day1.Task1()
		return
	}
	day1.Task2()
}

func taskDay2(task string) {
	if task == "1" {
		day2.Task1()
		return
	}
	day2.Task2()
}

func taskDay3(task string) {
	if task == "1" {
		day3.Task1()
		return
	}
	day3.Task2()
}

func taskDay4(task string) {
	if task == "1" {
		day4.Task1()
		return
	}
	day4.Task2()
}

func taskDay5(task string) {
	if task == "1" {
		day5.Task1()
		return
	}
	day5.Task2()
}

func taskDay6(task string) {
	if task == "1" {
		day6.Task1()
		return
	}
	day6.Task2()
}

func taskDay7(task string) {
	if task == "1" {
		day7.Task1()
		return
	}
	day7.Task2()
}

func taskDay8(task string) {
	if task == "1" {
		day8.Task1()
		return
	}
	day8.Task2()
}

func taskDay9(task string) {
	if task == "1" {
		day9.Task1()
		return
	}
	day9.Task2()
}
