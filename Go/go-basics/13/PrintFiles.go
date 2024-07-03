package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var Reset = "\033[0m"
var Green = "\033[32m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var Red = "\033[31m"

func main() {
	path := "C:/Users/Khaiarovir/go/src/Go/go-basics"
	PrintAllFiles(path)
	PrintAllFilesWithFilter(path, "12")
	PrintAllFilesWithFilterClosure(path, "12")
	PrintFilesWithFuncFilter(path, containsTwelve)
}

func PrintAllFiles(path string) {
	dir, err := os.ReadDir(path)
	if err != nil {
		log.Fatalf("Ошибка чтения директории %v", err)
	}
	for _, v := range dir {
		fmt.Println(Green + v.Name() + Reset)
		if v.IsDir() {
			name, _ := v.Info()
			PrintAllFiles(path + "/" + name.Name())
		}
	}
}
func PrintAllFilesWithFilter(path string, filter string) {
	dir, err := os.ReadDir(path)
	if err != nil {
		log.Fatalf("Ошибка чтения директории %v", err)
	}
	for _, v := range dir {
		filename := filepath.Join(path, v.Name())
		if v.IsDir() {
			PrintAllFilesWithFilter(filename, filter)
			continue
		}
		if strings.Contains(filename, filter) {
			fmt.Println(Cyan + v.Name() + Reset)
		}
	}
}
func PrintAllFilesWithFilterClosure(path string, filter string) {
	var walk func(string)
	walk = func(path string) {
		dir, err := os.ReadDir(path)
		if err != nil {
			log.Printf("Ошибка чтения директории %v", err)
		}
		for _, v := range dir {
			filename := filepath.Join(path, v.Name())

			if strings.Contains(filename, filter) {
				fmt.Println(Gray + filename + Reset)
			}
			if v.IsDir() {
				walk(filename)
			}
		}
	}
	walk(path)
}
func PrintFilesWithFuncFilter(path string, predicate func(string) bool) {
	var walk func(string)
	walk = func(path string) {
		dir, err := os.ReadDir(path)
		if err != nil {
			fmt.Printf("Ошибка чтения директории %v", err)
			return
		}
		for _, v := range dir {
			filename := filepath.Join(path, v.Name())
			if v.IsDir() {
				walk(filename)
			} else {
				if predicate(filename) {
					fmt.Println(Red + v.Name() + Reset)
				}
			}
		}
	}
	walk(path)
}
func containsTwelve(str string) bool {
	return strings.Contains(str, "12")
}
