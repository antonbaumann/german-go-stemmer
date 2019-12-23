package main

import (
	"german-go-stemmer/lookup"
	"regexp"
	"strings"
)

func prelude(word string) string {
	word = strings.TrimSpace(word)
	word = strings.ToLower(word)

	word = strings.Replace(word, "ß", "ss", -1)
	word = strings.Replace(word, "ä", "A", -1)
	word = strings.Replace(word, "ö", "O", -1)
	word = strings.Replace(word, "ü", "V", -1)

	for i := 2; i < len(word); i++ {
		prev := word[i-2]
		curr := word[i-1]
		next := word[i]
		if lookup.IsVowel(prev) && lookup.IsVowel(next) {
			if curr == 'u' {
				word = word[:i-1] + "U" + word[i:]
			}
			if curr == 'y' {
				word = word[:i-1] + "Y" + word[i:]
			}
		}
	}
	return word
}

func postlude(word string) string {
	word = strings.Replace(word, "A", "a", -1)
	word = strings.Replace(word, "O", "o", -1)
	word = strings.Replace(word, "V", "u", -1)
	word = strings.Replace(word, "Y", "y", -1)
	word = strings.Replace(word, "U", "u", -1)
	return word
}

func removeSuffixLst(word string, suffixes []string, r1 int) (string, bool) {
	for _, suff := range suffixes {
		if len(word) >= r1 && strings.HasSuffix(word[r1:], suff) {
			return strings.TrimSuffix(word, suff), true
		}
	}
	return word, false
}

func step1(word string, r1 int) string {
	a := []string{"em", "ern", "er"}
	word, changed := removeSuffixLst(word, a, r1)
	if changed {
		return word
	}

	b := []string{"e", "en", "es"}
	word, changed = removeSuffixLst(word, b, r1)
	if changed {
		if strings.HasSuffix(word, "niss") {
			word = strings.TrimSuffix(word, "s")
		}
		return word
	}

	if len(word) >= 2 && strings.HasSuffix(word[r1:], "s") && lookup.IsSEnding(word[len(word)-2]) {
		word = strings.TrimSuffix(word, "s")
	}

	return word
}

func step2(word string, r1 int) string {
	a := []string{"en", "er", "est"}
	word, changed := removeSuffixLst(word, a, r1)
	if changed {
		return word
	}

	if len(word) >= 6 && strings.HasSuffix(word, "st") && lookup.IsStEnding(word[len(word)-3]) {
		word = strings.TrimSuffix(word, "st")
	}
	return word
}

func step3(word string, r1, r2 int) string {
	a := []string{"end", "ung"}
	b := []string{"ig", "ik", "isch"}
	c := []string{"lich", "heit"}
	d := []string{"keit"}

	word, changed := removeSuffixLst(word, a, r2)
	if changed {
		if len(word) >= r2 && strings.HasSuffix(word[r2:], "ig") {
			if len(word) == 2 || word[len(word)-3] != 'e' {
				return strings.TrimSuffix(word, "ig")
			}
		}
		return word
	}

	for _, suff := range b {
		if len(word) >= r2 && strings.HasSuffix(word[r2:], suff) {
			if len(word) == len(suff) || word[len(word)-1-len(suff)] != 'e' {
				return strings.TrimSuffix(word, suff)
			}
		}
	}

	word, changed = removeSuffixLst(word, c, r2)
	if changed {
		word, _ = removeSuffixLst(word, []string{"er", "en"}, r1)
		return word
	}

	word, changed = removeSuffixLst(word, d, r2)
	if changed {
		word, _ = removeSuffixLst(word, []string{"lich", "ig"}, r2)
		return word
	}

	return word
}

func getRegion(word string, n int) int {
	if len(word) == 0 {
		return 0
	}

	for i := 1; i < len(word); i++ {
		if lookup.IsVowel(word[i-1]) && !lookup.IsVowel(word[i]) {
			if n == 0 {
				return i + 1
			} else {
				n -= 1
			}
		}
	}

	return len(word)
}

func getRegions(word string) (int, int) {
	r1 := getRegion(word, 0)
	r2 := getRegion(word, 1)

	if r1 < 3 && len(word) >= 3 {
		r1 = 3
	}

	return r1, r2
}

func stem(word string) string {
	word = prelude(word)
	r1, r2 := getRegions(word)
	word = step1(word, r1)
	word = step2(word, r1)
	word = step3(word, r1, r2)
	word = postlude(word)
	return word
}

func Stem(query string) string {
	r := regexp.MustCompile("[^\\s]+")
	words := r.FindAllString(query, -1)
	newQuery := make([]string, 0, len(words))
	for _, word := range words {
		if !lookup.IsStopWord(word) {
			res := stem(word)
			newQuery = append(newQuery, res)
		}
	}
	return strings.Join(newQuery[:], " ")
}

func StemWord(word string) string {
	return stem(word)
}

func StemWords(words []string) []string {
	stemmed := make([]string, len(words))
	for i, v := range words {
		stemmed[i] = stem(v)
	}
	return stemmed
}
