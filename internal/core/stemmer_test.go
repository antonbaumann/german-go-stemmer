package core

import (
	"fmt"
	"strings"
	"testing"

	"github.com/antonbaumann/german-go-stemmer/internal/util"
)

type testSet struct {
	value    string
	expected string
}

type regionsTestSet struct {
	value string
	r1    int
	r2    int
}

const errorMsg = "test failed.\nexpected:\t %v\ngot:\t\t %v"

func TestStemWord(t *testing.T) {
	words, err := util.ReadWordList("../../testdata/voc.txt")
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	expected, err := util.ReadWordList("../../testdata/stemmed.txt")
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
		result := Stem(test.value)
		if !strings.EqualFold(test.expected, result) {
			t.Errorf(errorMsg, test.expected, result)
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
		r1, _ := getRegions([]rune(test.value))
		result := step1([]rune(test.value), r1)
		if !strings.EqualFold(test.expected, string(result)) {
			t.Errorf(errorMsg, test.expected, result)
		}
	}
}

func TestStep2(t *testing.T) {
	tests := []testSet{
		{"derbst", "derb"},
	}
	for _, test := range tests {
		r1, _ := getRegions([]rune(test.value))
		result := step2([]rune(test.value), r1)
		if !strings.EqualFold(test.expected, string(result)) {
			t.Errorf(errorMsg, test.expected, result)
		}
	}
}

func TestGetRegions(t *testing.T) {
	tests := []regionsTestSet{
		{
			value: "baum",
			r1:    4,
			r2:    4,
		},
		{
			value: "",
			r1:    0,
			r2:    0,
		},
		{
			value: "alle",
			r1:    3,
			r2:    4,
		},
		{
			value: "lllläll",
			r1:    6,
			r2:    7,
		},
		{
			value: "aladin",
			r1:    3,
			r2:    4,
		},
	}

	for _, test := range tests {
		r1, r2 := getRegions([]rune(test.value))
		if r1 != test.r1 || r2 != test.r2 {
			fmt.Printf("tested %v\n", test.value)
		}
		if r1 != test.r1 {
			t.Errorf(errorMsg, test.r1, r1)
		}
		if r2 != test.r2 {
			t.Errorf(errorMsg, test.r2, r2)
		}
	}
}
