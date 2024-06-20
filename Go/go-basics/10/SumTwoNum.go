package main

import (
	"fmt"
)

func main() {
	arr := []int{3, 2, 5, 1, 7, 9}
	num := 14

	one, two := sumTwoNum(arr, num)
	fmt.Printf("Из массива, данных %v \n Сумма %d от двух чисел %d и %d \n Индексы равны %d, %d \n",
		arr, num, arr[one], arr[two], one, two)
	arr2 := sumTwo(arr, num)
	fmt.Println(arr2)
}
func sumTwoNum(arr []int, num int) (int, int) {
	for i := 0; i < len(arr); i++ {
		numTmp := num - arr[i]
		for j := 0; j < len(arr); j++ {
			if numTmp <= 0 || i == j {
				continue
			} else if numTmp == arr[j] {
				return i, j
			}
		}
	}
	return 0, 0
}
func sumTwo(arr []int, num int) []int {
	mapa := make(map[int]int)
	for k, v := range arr {
		i, ok := mapa[num-v]
		if ok {
			return []int{i, k}
		}
		mapa[v] = k
	}
	return nil
}
