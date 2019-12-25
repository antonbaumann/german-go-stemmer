package stemmer

import (
	"strings"

	"github.com/antonbaumann/german-go-stemmer/internal/lookup"
	"github.com/antonbaumann/german-go-stemmer/internal/util"
)

func replaceWithArray(a []rune, i int, v []rune) []rune {
	return append(a[:i], append(v, a[i+1:]...)...)
}

func prelude(word string) []rune {
	word = strings.TrimSpace(word)
	// convert word to rune array so multi-byte runes are easier to handle
	// will be converted back to string in postlude
	runes := []rune(word)

	for i := 0; i < len(runes); i++ {
		switch runes[i] {
		case 'ß':
			runes = replaceWithArray(runes, i, []rune{'s', 's'})
		case 'ä':
			runes[i] = 'A'
		case 'ö':
			runes[i] = 'O'
		case 'ü':
			runes[i] = 'V'
		case 'u':
			if i > 0 && i < len(runes)-1 && lookup.IsVowel(runes[i-1]) && lookup.IsVowel(runes[i+1]) {
				runes[i] = 'U'
			}
		case 'y':
			if i > 0 && i < len(runes)-1 && lookup.IsVowel(runes[i-1]) && lookup.IsVowel(runes[i+1]) {
				runes[i] = 'Y'
			}
		}
	}

	return runes
}

func postlude(runes []rune) string {
	for i := 0; i < len(runes); i++ {
		switch runes[i] {
		case 'A':
			runes[i] = 'a'
		case 'O':
			runes[i] = 'o'
		case 'V':
			runes[i] = 'u'
		case 'Y':
			runes[i] = 'y'
		case 'U':
			runes[i] = 'u'
		}
	}
	return string(runes)
}

func removeSuffixIfInList(runes []rune, suffixes [][]rune, r1 int) ([]rune, bool) {
	for _, suff := range suffixes {
		if len(runes) >= r1 && util.HasSuffix(runes[r1:], suff) {
			return runes[:len(runes)-len(suff)], true
		}
	}
	return runes, false
}

func step1(runes []rune, r1 int) []rune {
	a := [][]rune{
		{'e', 'm'},
		{'e', 'r', 'n'},
		{'e', 'r'},
	}
	b := [][]rune{
		{'e'},
		{'e', 'n'},
		{'e', 's'},
	}

	runes, changed := removeSuffixIfInList(runes, a, r1)
	if changed {
		return runes
	}

	runes, changed = removeSuffixIfInList(runes, b, r1)
	if changed {
		if util.HasSuffix(runes, []rune{'n', 'i', 's', 's'}) {
			runes = runes[:len(runes)-1]
		}
		return runes
	}

	suff := []rune{'s'}
	if len(runes) >= 2 && util.HasSuffix(runes[r1:], suff) && lookup.IsSEnding(runes[len(runes)-len(suff)-1]) {
		runes = runes[:len(runes)-len(suff)]
	}

	return runes
}

func step2(runes []rune, r1 int) []rune {
	a := [][]rune{
		{'e', 'n'},
		{'e', 'r'},
		{'e', 's', 't'},
	}
	runes, changed := removeSuffixIfInList(runes, a, r1)
	if changed {
		return runes
	}

	suff := []rune{'s', 't'}
	if len(runes) >= 6 && util.HasSuffix(runes, suff) && lookup.IsStEnding(runes[len(runes)-3]) {
		runes = runes[:len(runes)-len(suff)]
	}
	return runes
}

func step3(runes []rune, r1, r2 int) []rune {
	a := [][]rune{
		{'e', 'n', 'd'},
		{'u', 'n', 'g'},
	}
	b := [][]rune{
		{'i', 'g'},
		{'i', 'k'},
		{'i', 's', 'c', 'h'},
	}
	c := [][]rune{
		{'l', 'i', 'c', 'h'},
		{'h', 'e', 'i', 't'},
	}
	d := [][]rune{
		{'k', 'e', 'i', 't'},
	}

	runes, changed := removeSuffixIfInList(runes, a, r2)
	if changed {
		suff := []rune{'i', 'g'}
		if len(runes) >= r2 && util.HasSuffix(runes[r2:], suff) {
			if len(runes) == len(suff) || runes[len(runes)-len(suff)-1] != 'e' {
				return runes[:len(runes)-len(suff)]
			}
		}
		return runes
	}

	for _, suff := range b {
		if len(runes) >= r2 && util.HasSuffix(runes[r2:], suff) {
			if len(runes) == len(suff) || runes[len(runes)-len(suff)-1] != 'e' {
				return runes[:len(runes)-len(suff)]
			}
		}
	}

	runes, changed = removeSuffixIfInList(runes, c, r2)
	if changed {
		suffixes := [][]rune{{'e', 'r'}, {'e', 'n'}}
		runes, _ = removeSuffixIfInList(runes, suffixes, r1)
		return runes
	}

	runes, changed = removeSuffixIfInList(runes, d, r2)
	if changed {
		suffixes := [][]rune{{'l', 'i', 'c', 'h'}, {'i', 'g'}}
		runes, _ = removeSuffixIfInList(runes, suffixes, r2)
		return runes
	}

	return runes
}

func getRegion(runes []rune, n int) int {
	if len(runes) == 0 {
		return 0
	}

	for i := 0; i < len(runes); i++ {
		if i >= 1 && lookup.IsVowel(runes[i-1]) && !lookup.IsVowel(runes[i]) {
			if n == 0 {
				return i + 1
			}
			n--
		}
	}

	return len(runes)
}

func getRegions(runes []rune) (int, int) {
	r1 := getRegion(runes, 0)
	r2 := getRegion(runes, 1)

	if r1 < 3 && len(runes) >= 3 {
		r1 = 3
	}

	return r1, r2
}

func Stem(word string) string {
	runes := prelude(word)
	r1, r2 := getRegions(runes)
	runes = step1(runes, r1)
	runes = step2(runes, r1)
	runes = step3(runes, r1, r2)
	word = postlude(runes)
	return word
}
