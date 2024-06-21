package main

import "fmt"

func main() {
	input := []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	}
	fmt.Println(input)
	fmt.Println(removeDublicates(input))
}

func removeDublicates(input []string) []string {
	output := make([]string, 0)
	mapa := make(map[string]int)
	for k, v := range input {
		_, ok := mapa[v]
		if !ok {
			output = append(output, v)
		}
		mapa[v] = k
	}
	return output
}
