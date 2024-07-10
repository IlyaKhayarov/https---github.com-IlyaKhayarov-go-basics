package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	openFile()
	fmt.Printf("unintuitive = %v", unintuitive())
	fmt.Printf("\nintuitive = %v", intuitive())
}
func unintuitive() (value string) {
	defer func() { value = "На самом деле" }() // круглые скобки в конце означают, что функция вызывается
	return "Казалось бы"
}
func intuitive() string {
	value := "Казалось бы"
	defer func() { value = "На самом деле" }()
	return value
}

func openFile() {
	defer metricTime(time.Now())

	file, err := os.OpenFile("text.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	RaiseErr(err)
	defer file.Close()

	_, err = file.WriteString("1")
	RaiseErr(err)

	dat, err := os.ReadFile(file.Name())
	RaiseErr(err)
	fmt.Println(string(dat))
}

func RaiseErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func metricTime(start time.Time) {
	fmt.Println(time.Now().Sub(start))
}
