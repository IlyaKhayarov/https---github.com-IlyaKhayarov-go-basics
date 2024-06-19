package main

import (
	"fmt"
	"go/go-basics/9/price"
)

func main() {
	charCount("Πολύ λίγα πράγματα συμβαίνουν στο σωστό χρόνο, και τα υπόλοιπα δεν συμβαίνουν καθόλου")
	price.DelicaciesList()
}

func charCount(sentence string) {
	fmt.Printf("Start count of char..\n")
	frequency := make(map[rune]int)

	for _, k := range sentence {
		frequency[k]++
	}
	for k, v := range frequency {
		fmt.Printf("Значение %c встречается %d раз \n", k, v)
	}
}
