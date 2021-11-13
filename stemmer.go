package stemmer

import (
	"regexp"
	"strings"

	"github.com/antonbaumann/german-go-stemmer/internal/core"
	"github.com/antonbaumann/german-go-stemmer/internal/lookup"
)

// Stem stems a raw query and removes stop-words
func Stem(query string) string {
	r := regexp.MustCompile("[^\\s]+")
	words := r.FindAllString(query, -1)
	newQuery := make([]string, 0, len(words))
	for _, word := range words {
		if !lookup.IsStopWord(word) {
			res := core.Stem(word)
			newQuery = append(newQuery, res)
		}
	}
	return strings.Join(newQuery, " ")
}

// StemWord stems a single keyword
func StemWord(word string) string {
	return core.Stem(word)
}

// StemWords stems a list of keywords
func StemWords(words []string) []string {
	stemmed := make([]string, len(words))
	for i, v := range words {
		stemmed[i] = core.Stem(v)
	}
	return stemmed
}
