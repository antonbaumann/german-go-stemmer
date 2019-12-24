package stemmer

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

type testSet struct {
	value    string
	expected string
}

func TestStemQuery(t *testing.T) {
	query := "wie wird das wetter  \t morgen in \nmünchen"
	expected := "wett morg munch"
	result := Stem(query)

	if !strings.EqualFold(expected, result) {
		t.Errorf("test failed.\nexpected:\t %v\ngot:\t\t %v", expected, result)
	}
}

func TestStemWords(t *testing.T) {
	words, err := readWordList("testdata/voc.txt")
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	expected, err := readWordList("testdata/stemmed.txt")
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if len(words) != len(expected) {
		t.Errorf("wordlist and result list should have same size")
		t.Fail()
	}

	result := StemWords(words)

	if !slicesEqual(expected, result) {
		t.Errorf("test failed.\nexpected:\t %v\ngot:\t\t %v", expected, result)
	}
}

func TestStemWord(t *testing.T) {
	words, err := readWordList("testdata/voc.txt")
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	expected, err := readWordList("testdata/stemmed.txt")
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if len(words) != len(expected) {
		t.Errorf("wordlist and result list should have same size")
		t.Fail()
	}

	tests := make([]*testSet, 0, len(words))
	for i, v := range words {
		tests = append(tests, &testSet{value: v, expected: expected[i]})
	}

	for _, test := range tests {
		result := StemWord(test.value)
		if !strings.EqualFold(test.expected, result) {
			t.Errorf("test failed.\nexpected:\t %v\ngot:\t\t %v", test.expected, result)
		}
	}
}

func TestStep1(t *testing.T) {
	tests := []testSet{
		{"Ackern", "Ack"},
		{"ackers", "acker"},
		{"armes", "arm"},
		{"bedUrfnissen", "bedUrfnis"},
	}
	for _, test := range tests {
		r1, _ := getRegions(test.value)
		result := step1(test.value, r1)
		if !strings.EqualFold(test.expected, result) {
			t.Errorf("test failed.\nexpected:\t %v\ngot:\t\t %v", test.expected, result)
		}
	}
}

func TestStep2(t *testing.T) {
	tests := []testSet{
		{"derbst", "derb"},
	}
	for _, test := range tests {
		r1, _ := getRegions(test.value)
		result := step2(test.value, r1)
		if !strings.EqualFold(test.expected, result) {
			t.Errorf("test failed.\nexpected:\t %v\ngot:\t\t %v", test.expected, result)
		}
	}
}

func TestGetRegionEmptyWord(t *testing.T) {
	word := ""
	r1, r2 := getRegions(word)
	if r1 != 0 {
		t.Errorf("test failed.\nexpected:\t %v\ngot:\t\t %v", 0, r1)
	}
	if r2 != 0 {
		t.Errorf("test failed.\nexpected:\t %v\ngot:\t\t %v", 0, r2)
	}
}

func readWordList(filePath string) ([]string, error) {
	words := make([]string, 0)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	return words, nil
}

func slicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if !strings.EqualFold(v, b[i]) {
			return false
		}
	}
	return true
}
