package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var found = 1

func main() {
	query := os.Args[2]
	root := os.Args[1]
	filepath.Walk(root, func(path string, file os.FileInfo, err error) error {
		if !file.IsDir() {
			file, err := os.Open(path)
			defer file.Close()
			if err != nil {
				return nil
			}
			scanner := bufio.NewScanner(file)
			for i := 1; scanner.Scan(); i++ {
				if strings.Contains(scanner.Text(), query) {
					found = 0
					fmt.Printf("%s/%s:%d: %s\n", root, path, i, scanner.Text())
				}
			}
		}
		return nil
	})
}

//PS C:\Golang_work\src\GoCourse\homework-5\gogrep> go run .\gogrep.go C:\Golang_work\src\GoCourse\homework-5 тест
//C:\Golang_work\src\GoCourse\homework-5/C:\Golang_work\src\GoCourse\homework-5\gofc\test.txt:2: тест
//C:\Golang_work\src\GoCourse\homework-5/C:\Golang_work\src\GoCourse\homework-5\homework-5.go:104: //тест
//C:\Golang_work\src\GoCourse\homework-5/C:\Golang_work\src\GoCourse\homework-5\test.txt:2: тест
