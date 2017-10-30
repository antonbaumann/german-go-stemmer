package main

import (
	"os"
	"log"
	"bufio"
	"strings"
)

func ScanFile(path string) []string {
	var tmp []string
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tmp = append(tmp, strings.Split(scanner.Text(), " ")[0])
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return tmp
}
