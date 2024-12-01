package main

import (
	"os"

	"github.com/zanBizjak/advent-of-code-24/day1"
)

func main() {
	day := os.Args[1]
	task := os.Args[2]

	switch day {
	case "1":
		taskDay1(task)
	}

}

func taskDay1(task string) {
	if task == "1" {
		day1.Task1()
		return
	}
	day1.Task2()

}
