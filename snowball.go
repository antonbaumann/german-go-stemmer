package stemmer

import (
	"regexp"
	"strings"

	"github.com/antonbaumann/german-go-stemmer/internal/lookup"
	"github.com/antonbaumann/german-go-stemmer/internal/stemmer"
)

// Stem stems a raw query and removes stop-words
func Stem(query string) string {
	r := regexp.MustCompile("[^\\s]+")
	words := r.FindAllString(query, -1)
	newQuery := make([]string, 0, len(words))
	for _, word := range words {
		if !lookup.IsStopWord(word) {
			res := stemmer.Stem(word)
			newQuery = append(newQuery, res)
		}
	}
	return strings.Join(newQuery[:], " ")
}

// StemWord stems a single keyword
func StemWord(word string) string {
	return stemmer.Stem(word)
}

// StemWords stems a list of keywords
func StemWords(words []string) []string {
	stemmed := make([]string, len(words))
	for i, v := range words {
		stemmed[i] = stemmer.Stem(v)
	}
	return stemmed
}
