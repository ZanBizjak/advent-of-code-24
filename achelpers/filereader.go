package achelpers

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func StringReadFile(file string) string {
	dat, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	return string(dat)
}

func IntReadTwoColumns(filestring string, splitter string) ([]int, []int) {
	file, err := os.Open(filestring)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	col1 := make([]int, 1000)
	col2 := make([]int, 1000)

	for scanner.Scan() {
		row := scanner.Text()
		fmt.Println(row)
		firstNum, err := strconv.Atoi(strings.Split(row, splitter)[0])
		if err != nil {
			log.Fatal(err)
		}
		secondNum, err := strconv.Atoi(strings.Split(row, splitter)[1])
		if err != nil {
			log.Fatal(err)
		}
		col1 = append(col1, firstNum)
		col2 = append(col2, secondNum)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return col1, col2
}
