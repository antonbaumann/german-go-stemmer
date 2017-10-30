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

func TestStem2(t *testing.T) {

}

func TestStemQuery(t *testing.T) {
	queries := []string{
		"wie kann ich schnell abnehmen",
	}
	for _, query := range queries {
		fmt.Println(StemQuery(query))
	}
}

//func TestStem2(t *testing.T) {
//	word := "kätzchens"
//	var err error
//	word0, err := step0(word)
//	if err != nil {log.Println(err)}
//	fmt.Println(0, word0)
//	word1, err := step1(word0)
//	if err != nil {log.Println(err)}
//	fmt.Println(1, word1)
//	word2 := step2(word1)
//	fmt.Println(2, word2)
//	word3 := step3(word2)
//	fmt.Println(3, word3)
//	word4 := step4(word3)
//	fmt.Println(4, word4)
//	fmt.Println(word4)
//}

