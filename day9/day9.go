package day9

import (
	"fmt"

	"github.com/zanBizjak/advent-of-code-24/achelpers"
)

func Task1() {
	disk := achelpers.ReadRows("day9.txt")[0]
	var idExpanded []int64

	action := true
	var id int64 = 0
	for _, info := range disk {
		if action {
			for i := 0; i < int(info-'0'); i++ {
				idExpanded = append(idExpanded, id)
			}
			id++
		} else {
			for i := 0; i < int(info-'0'); i++ {
				idExpanded = append(idExpanded, -1)
			}
		}
		action = !action
	}

	i := 0
	j := len(idExpanded) - 1
	for i < j {
		for idExpanded[j] == -1 {
			j--
		}
		for idExpanded[i] != -1 {
			i++
		}
		if i < j {

			a := idExpanded[i]
			idExpanded[i] = idExpanded[j]
			idExpanded[j] = a
		}

	}

	var res int64 = 0
	var k int64 = 0
	for ; idExpanded[k] != -1; k++ {
		res += int64(idExpanded[k]) * k

	}

	fmt.Println(res)

}

func Task2() {
}
