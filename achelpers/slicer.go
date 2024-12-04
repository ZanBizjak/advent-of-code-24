package achelpers

func IntRemoveIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func RuneCopySlice(s []rune) []rune {
	ret := make([]rune, 0)
	return append(ret, s...)
}
