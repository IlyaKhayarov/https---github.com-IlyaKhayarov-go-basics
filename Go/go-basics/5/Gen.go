package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	var year int
	fmt.Println("Введите год рождения")
	for {
		y, bl := check(strconv.Atoi(scan()))
		if bl == true {
			year = y
			break
		}
	}
	switch {
	case year > 1900 && year < 1946:
		fmt.Println("Мое почтение!")
	case year >= 1946 && year <= 1964:
		fmt.Println("Привет, бумер!")
	case year >= 1965 && year <= 1980:
		fmt.Println("Привет, представитель X!")
	case year >= 1981 && year <= 1996:
		fmt.Println("Привет, миллениал!")
	case year >= 1997 && year <= 2012:
		fmt.Println("Привет, зумер!")
	case year >= 2013 && year <= time.Now().Year():
		fmt.Println("Привет, альфа!")
	default:
		fmt.Println("Привет, еще не родившимся!")
	}
}
func check(y int, err error) (int, bool) {
	if err != nil {
		fmt.Println("Введите цифрами год рождения!")
	} else {
		if y < 1900 {
			fmt.Println("Введите верный год!")
		} else {
			return y, true
		}
	}
	return y, false
}
func scan() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
