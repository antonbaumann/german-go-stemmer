package stemmer

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type testSet struct {
	value    string
	expected string
}

func TestStemQuery(t *testing.T) {
	query := "wie wird das wetter  \t morgen in \nmünchen"
	result := Stem(query)
	assert.Equal(t, "wett morg munch", result)
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

	assert.Equal(t, expected, StemWords(words))
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
		result := stem(test.value)
		assert.Equal(t, test.expected, result)
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
		assert.Equal(t, test.expected, result)
	}
}

func TestStep2(t *testing.T) {
	tests := []testSet{
		{"derbst", "derb"},
	}
	for _, test := range tests {
		r1, _ := getRegions(test.value)
		result := step2(test.value, r1)
		assert.Equal(t, test.expected, result)
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
