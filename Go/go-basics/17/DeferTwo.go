package main

import "fmt"

var Global = 5

func main() {
	fmt.Printf("Global value = %v\n", Global)
	PrintGlobalVal()
	fmt.Printf("Global value = %v\n", Global)
}

func PrintGlobalVal() {
	defer func(val int) {
		Global = val
	}(Global)
	Global = 2
	fmt.Printf("Global value = %v\n", Global)
}
