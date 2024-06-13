package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i <= 100; i++ {
		str := "FizzBuzz"
		switch {
		case (i%5 == 0 && i%3 == 0):
		case i%3 == 0:
			str = "Fizz"
		case i%5 == 0:
			str = "Buzz"
		default:
			str = strconv.Itoa(i)
		}
		fmt.Println(str)
	}
}
