package main

import (
	"testing"
	"os"
	"bufio"
	"log"
	"fmt"
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
		tmp = append(tmp,scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return tmp
}


func TestStem(t *testing.T) {
	var words 				= ScanFile("txt/words_to_stem.txt")
	var correctly_stemmed 	= ScanFile("txt/correctly_stemmed_words.txt")
	var stemmed 			[]string

	for _, w := range words {
		//fmt.Println("testing " + w)
		st, err := Stem(w)
		if err != nil {
			t.Error(err)
		}
		stemmed = append(stemmed, st)
	}

	for i, w := range words {
		fmt.Printf("%-30s %-30s %-3t\n", w, stemmed[i], strings.EqualFold(correctly_stemmed[i], stemmed[i]))
		//fmt.Println(w, stemmed[i], strings.EqualFold(correctly_stemmed[i], stemmed[i]))
		if !strings.EqualFold(correctly_stemmed[i], stemmed[i]) {
			fmt.Println("! ", correctly_stemmed[i])
		}
	}
}

