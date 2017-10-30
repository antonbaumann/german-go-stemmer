package main

import (
	"testing"
	"fmt"
	"strings"
)


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
		fmt.Printf("%-25s %-20s %-3t\n", w, stemmed[i], strings.EqualFold(correctly_stemmed[i], stemmed[i]))
		if !strings.EqualFold(correctly_stemmed[i], stemmed[i]) {
			fmt.Println("! ", correctly_stemmed[i])
		}
	}
}


func TestStemQuery(t *testing.T) {
	fmt.Println("\n\n", "TEST StemQuery()", "\n")
	queries := []string {
		"ich will schnell abnehmen",
		"wie kann ich schnell abnehmen",
		"welcher film läuft heute im kino",
		"was ist die beste pizzaria in münchen",
		"wie kann ich schnell abnehmen ohne sport zu machen",
	}
	for _, query := range queries {
		fmt.Println(StemQuery(query))
	}
}
