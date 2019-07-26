package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestStem(t *testing.T) {
	words := ScanFile("txt/words_to_stem.txt")
	correctlyStemmed := ScanFile("txt/correctly_stemmed_words.txt")
	var stemmed []string

	for _, w := range words {
		st, err := Stem(w)
		if err != nil {
			t.Error(err)
		}
		stemmed = append(stemmed, st)
	}

	for i, w := range words {
		fmt.Printf("%-25s %-20s %-3t\n", w, stemmed[i], strings.EqualFold(correctlyStemmed[i], stemmed[i]))
		assert.Equal(t, correctlyStemmed[i], stemmed[i])
	}
}

func TestStemQuery(t *testing.T) {
	query := "wie kann ich schnell abnehmen"
	result, err := StemQuery(query)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "schnell abnehm", result)
}