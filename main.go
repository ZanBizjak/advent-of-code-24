package main

import (
	"os"

	"github.com/zanBizjak/advent-of-code-24/day1"
	"github.com/zanBizjak/advent-of-code-24/day2"
	"github.com/zanBizjak/advent-of-code-24/day3"
	"github.com/zanBizjak/advent-of-code-24/day4"
	"github.com/zanBizjak/advent-of-code-24/day5"
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
