package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var Green = "\033[32m"
var Cyan = "\033[36m"

func main() {
	PrintAllFiles("C:/Users/Khaiarovir/go/src/Go/go-basics")
	PrintAllFilesWithFilter("C:/Users/Khaiarovir/go/src/Go/go-basics", "12")
}

func PrintAllFiles(path string) {
	dir, err := os.ReadDir(path)
	if err != nil {
		log.Fatalf("Ошибка чтения директории %v", err)
	}
	for _, v := range dir {
		fmt.Println(v.Name() + Green)
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
			fmt.Println(v.Name() + Cyan)
		}
	}
}
