package main

import (
	"errors"
	"fmt"
)

func main() {
	startInt(100)
	startStr("The quick brown 狐 jumped over the lazy 犬")
}
func startInt(count int) {
	slice := make([]int, 0)

	fillSlice(&slice, count)
	err := getFirstLast(&slice, 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(slice)
	sliceReverseInt(&slice)
	fmt.Println(slice)
}
func startStr(input string) {
	slice := make([]rune, len(input))

	for i, val := range input {
		slice[i] = val
	}
	fmt.Println(string(slice))

	sliceReverseRune(&slice)
	fmt.Println(string(slice))
}
func fillSlice(sl *[]int, count int) {
	for i := 1; i <= count; i++ {
		*sl = append(*sl, i)
	}
}
func sliceReverseInt(sl *[]int) {
	dim := len(*sl)
	for i := range (*sl)[:dim/2] {
		(*sl)[i], (*sl)[dim-i-1] = (*sl)[dim-i-1], (*sl)[i]
	}
}
func sliceReverseRune(sl *[]rune) {
	dim := len(*sl)
	for i := range (*sl)[:dim/2] {
		(*sl)[i], (*sl)[dim-i-1] = (*sl)[dim-i-1], (*sl)[i]
	}
}
func getFirstLast(sl *[]int, i int) error {
	dim := len(*sl)
	if dim < i {
		return errors.New("Create slice bigger!")
	} else {
		*sl = append((*sl)[:i], (*sl)[len(*sl)-i:]...)
	}
	return nil
}
