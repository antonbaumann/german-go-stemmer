package stemmer

import (
	"strings"
	"testing"
)

type testSet struct {
	value    string
	expected string
}

func BenchmarkStemWord(b *testing.B) {
	words, err := readWordList("testdata/voc.txt")
	if err != nil {
		b.Error(err)
		b.Fail()
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = StemWord(words[i%len(words)])
	}
	b.StopTimer()
}

func BenchmarkStemQuery(b *testing.B) {
	words, err := readWordList("testdata/voc.txt")
	if err != nil {
		b.Error(err)
		b.Fail()
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = Stem(words[i%len(words)])
	}
	b.StopTimer()
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
