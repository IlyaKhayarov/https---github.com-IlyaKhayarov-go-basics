package price

import (
	"fmt"
	"strings"
	"unicode"
)

func DelicaciesList() {
	fmt.Printf("Delicacies list..\n")
	price := map[string]int{"ХЛЕБ": 50, "Молоко": 100, "масло": 200, "Колбаса": 500,
		"Соль": 20, "Огурцы": 200, "Сыр": 600, "Ветчина": 700, "буженина": 900,
		"Помидоры": 250, "Рыба": 300, "Хамон": 1500}
	getProductsGreaterSum(price, 500)
	getSumOfOrder(price, []string{"хлеб", "Буженина", "сыр", "огурцы", "МАСЛО"})
}
func getProductsGreaterSum(price map[string]int, dim int) {
	fmt.Printf("List of products greater then sum %d:\n", dim)
	for k, v := range price {
		if v > dim {
			fmt.Printf("Этот продукт %s стоит более %d рублей\n", k, dim)
		}
	}
}
func getSumOfOrder(price map[string]int, order []string) {
	sum := 0
	for _, v := range order {
		sum += getValueFromMap(price, v)
	}
	fmt.Printf("Стоимость заказа %s %d рублей\n", order, sum)
}
func getValueFromMap(price map[string]int, key string) int {
	pr, t := price[key]
	if !t {
		runes := []rune(key)
		runes[0] = unicode.ToUpper(runes[0])
		pr, t = price[string(runes)]
	}
	if !t {
		pr, t = price[strings.ToUpper(key)]
	}
	if !t {
		pr = price[strings.ToLower(key)]
	}
	return pr
}
