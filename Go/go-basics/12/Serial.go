package main

import (
	"encoding/json"
	"fmt"
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
	j, err := json.Marshal(p)
	if err != nil {
		log.Fatalln(`Ошибка сериализации`)
	} else {
		fmt.Println(string(j))
	}
}
