package main

import (
	"fmt"
	"slices"
	"time"
)

func ReverseNumbers(numbers []int) []int {
	lenthSlice := len(numbers)
	newSlice := make([]int, lenthSlice)
	for i := 0; i < lenthSlice; i++ {
		lenthSlice -= 1
		newSlice[i], newSlice[lenthSlice] = numbers[lenthSlice], numbers[i]
	}
	return newSlice
}

func main() {
	generalSlice := []int{}
	for i := 0; i < 11; i++ {
		generalSlice = append(generalSlice, i)
	}
	//fmt.Println(ReverseNumbers(generalSlice))
	fmt.Printf("speed = %d\n", speed(generalSlice))
	fmt.Printf("speed2 = %d\n", speed2(generalSlice))
}

func speed(sl []int) int64 {
	before := time.Now()
	fmt.Println(ReverseNumbers(sl))
	//fmt.Println(sl2)
	after := time.Now()
	//fmt.Println(int64(after.Sub(before)))
	return after.Sub(before).Nanoseconds()
}
func speed2(sl []int) int64 {
	before := time.Now()
	slices.Reverse(sl)
	//fmt.Println(sl)
	after := time.Now()
	return after.Sub(before).Nanoseconds()
}
