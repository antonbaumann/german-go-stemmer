package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

// äöü -> AOU so they can be represented as ascii chars
const vowels = "aeiouyAOV"
const sEnding = "bdfghklmnrt"
const stEnding = "bdfghklmnt"

var p1, p2 int

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

func step0(word string) (string, error) {
	word = strings.TrimSpace(word)
	word = strings.ToLower(word)

	word = strings.Replace(word, "ß", "ss", -1)
	word = strings.Replace(word, "ä", "A", -1)
	word = strings.Replace(word, "ö", "O", -1)
	word = strings.Replace(word, "ü", "V", -1)

	// if u,y between vowels -> ToUpper()
	for i := 2; i < len(word); i++ {
		prev := word[i-2]
		curr := word[i-1]
		next := word[i]
		if strings.Contains(vowels, string(prev)) && strings.Contains(vowels, string(next)) {
			if curr == 'u' {
				word = word[:i-1] + "U" + word[i:]
			}
			if curr == 'y' {
				word = word[:i-1] + "Y" + word[i:]
			}
		}
	}
	word = strings.Replace(word, "ß", "ss", -1)

	var err error
	p1, p2, err = getR(word)
	if err != nil {
		return "", err
	}

	return word, nil
}

func step1(word string) (string, error) {
	R1 := ""
	if p1 <= len(word) {
		R1 = word[p1:]
	}

	a := []string{"em", "ern", "er"}
	b := []string{"e", "en", "es"}

	for _, suffix := range a {
		if strings.HasSuffix(R1, suffix) {
			word = word[:len(word)-len(suffix)]
			return word, nil
		}
	}
	for _, suffix := range b {
		if strings.HasSuffix(R1, suffix) {
			word = word[:len(word)-len(suffix)]
			if strings.HasSuffix(word, "niss") {
				word = word[:len(word)-1]
			}
			return word, nil
		}
	}
	if strings.HasSuffix(word, "s") {
		rune, err := RuneAt(word, len(word)-2)
		if err != nil {
			return "", err
		}
		if strings.Contains(sEnding, string(rune)) {
			word = word[:len(word)-1]
		}
	}

	return word, nil
}

func step2(word string) string {
	R1 := ""
	if p1 <= len(word) {
		R1 = word[p1:]
	}

	a := []string{"an", "er", "est"}

	for _, suffix := range a {
		if strings.HasSuffix(R1, suffix) {
			word = word[:len(word)-len(suffix)]
			return word
		}
	}

	if strings.HasSuffix(R1, "st") && len(word) >= 6 {
		if strings.Contains(stEnding, string(word[len(word)-3])) {
			word = word[:len(word)-2]
		}
	}
	return word
}

func step3(word string) string {
	R1 := ""
	if p1 <= len(word) {
		R1 = word[p1:]
	}
	R2 := ""
	if p2 <= len(word) {
		R2 = word[p2:]
	}
	a := []string{"end", "ung"}
	b := []string{"ig", "ik", "isch"}
	c := []string{"lich", "heit"}
	d := []string{"keit"}

	for _, suffix := range a {
		// On website they use R2
		if strings.HasSuffix(R1, suffix) {
			word = word[:len(word)-len(suffix)]
			R1 = ""
			if p1 <= len(word) {
				R1 = word[p1:]
			}
			if strings.HasSuffix(R2, "ig") {
				i := len(word) - 3
				if i >= 0 && word[i] == 'e' {
					word = word[:len(word)-1]
				}
			}
			return word
		}
	}

	for _, suffix := range b {
		if strings.HasSuffix(R2, suffix) {
			i := len(word) - 3
			if i >= 0 && word[i] != 'e' {
				word = word[:len(word)-len(suffix)]
				return word
			}
		}
	}

	for _, suffix := range c {
		if strings.HasSuffix(R2, suffix) {
			word = word[:len(word)-len(suffix)]
			R2 = ""
			if p2 <= len(word) {
				R2 = word[p2:]
			}
			if strings.HasSuffix(R1, "er") || strings.HasSuffix(R1, "en") {
				word = word[:len(word)-2]
			}
			return word
		}
	}

	for _, suffix := range d {
		if strings.HasSuffix(R2, suffix) {
			word = word[:len(word)-len(suffix)]
			R2 = ""
			if p2 <= len(word) {
				R2 = word[p2:]
			}
			if strings.HasSuffix(R2, "lich") {
				word = word[:len(word)-4]
			} else if strings.HasSuffix(R2, "ig") {
				word = word[:len(word)-2]
			}
			return word
		}
	}

	return word
}

func step4(word string) string {
	word = strings.Replace(word, "A", "a", -1)
	word = strings.Replace(word, "O", "o", -1)
	word = strings.Replace(word, "V", "u", -1)
	word = strings.Replace(word, "Y", "y", -1)
	word = strings.Replace(word, "U", "u", -1)
	return word
}

func getR(word string) (int, int, error) {
	var p1, p2 int

	//for i := 3; i<len(word); i++ {
	//	prev, err := RuneAt(word, i-1)
	//	if err != nil {return -1, -1, err}
	//	pos, err := RuneAt(word, i)
	//	if err != nil {return -1, -1, err}
	//	if strings.ContainsRune(vowels, prev) && !strings.ContainsRune(vowels, pos) {
	//		p1 = i+1
	//		break
	//	}
	//}
	//if p1 == 0 || p1 >= len(word){
	for i := 1; i < len(word); i++ {
		prev, err := RuneAt(word, i-1)
		if err != nil {
			return -1, -1, err
		}
		pos, err := RuneAt(word, i)
		if err != nil {
			return -1, -1, err
		}
		if strings.ContainsRune(vowels, prev) && !strings.ContainsRune(vowels, pos) {
			p1 = i + 1
			break
		}
	}
	//}
	for i := p1 + 1; i < len(word); i++ {
		prev, err := RuneAt(word, i-1)
		if err != nil {
			return -1, -1, err
		}
		pos, err := RuneAt(word, i)
		if err != nil {
			return -1, -1, err
		}
		if strings.ContainsRune(vowels, prev) && !strings.ContainsRune(vowels, pos) {
			p2 = i + 1
			break
		}
	}
	if p2 == 0 {
		p2 = len(word)
	}
	if p1 == 0 {
		p1 = len(word)
	}

	return p1, p2, nil
}

func RuneAt(string string, int int) (rune, error) {
	if int >= len(string) {
		err := fmt.Errorf("index out of bound: %v (length is %v)", int, len(string))
		return -1, err
	}
	return []rune(string)[int], nil
}

func Stem(word string) (string, error) {
	var err error
	word, err = step0(word)
	if err != nil {
		return "", err
	}
	word, err = step1(word)
	if err != nil {
		return "", err
	}
	word = step2(word)
	word = step3(word)
	word = step4(word)
	return word, nil
}

func StemQuery(query string) (string, error) {
	stopWords := ScanFile("txt/stop_words_de.txt")
	sort.Strings(stopWords)
	r := regexp.MustCompile("[^\\s]+")
	words := r.FindAllString(query, -1)
	newQuery := ""
	for _, w := range words {
		if stopWords[sort.SearchStrings(stopWords, w)] != w {
			res, err := Stem(w)
			if err != nil {
				return "", err
			}
			newQuery += res + " "
		}
	}
	return strings.TrimSpace(newQuery), nil
}
