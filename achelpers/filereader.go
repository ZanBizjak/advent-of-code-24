package achelpers

import (
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func StringReadFile(file string) string {
	dat, err := os.ReadFile(file)
	check(err)

	return string(dat)
}
