package main

import (
	"encoding/json"
	"fmt"
	"go/go-basics/12/deserial"
	"log"
	"time"
)

type Person struct {
	Name        string    `json:"Имя"`
	Email       string    `json:"Почта"`
	DateOfBirth time.Time `json:"-"`
}

func main() {
	p := Person{Name: "Алекс", Email: "alex@yandex.ru"}

	fmt.Printf("Запуск сериализации %v\n", p)
	j, err := json.Marshal(p)
	if err != nil {
		log.Fatalln(`Ошибка сериализации\n`)
	} else {
		fmt.Printf("Результат: %v\n", string(j))
	}

	fmt.Printf("Запуск десериализации\n")
	r, err := deserial.ReadRespond(getJson())
	if err != nil {
		log.Fatalln("Ошибка десериализации\n")
	} else {
		fmt.Printf("Ответ: %+v\n", r)
	}
}
func getJson() string {
	return `{
		"header": {
			"code": 0,
			"message": "Message example"
		},
		"data": [{
			"type": "user",
			"id": 100,
			"attributes": {
				"email": "bob@yandex.ru",
				"article_ids": [10, 11, 12]
			}
		}]
	}`
}
